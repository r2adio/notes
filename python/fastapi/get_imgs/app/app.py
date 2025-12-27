import os
import shutil
import tempfile
import uuid
from contextlib import asynccontextmanager

from click.types import convert_type
from fastapi import Depends, FastAPI, File, Form, HTTPException, UploadFile

# from imagekitio.models.UploadFileRequestOptions import UploadFileRequestOptions
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from app.db import Post, create_db, get_async_session
from app.images import imagekit
from app.schemas import PostCreate, PostResponse


@asynccontextmanager
async def lifespan(app: FastAPI):
    await create_db()
    yield


# create the FastAPI app, and add the lifespan context manager
app = FastAPI(lifespan=lifespan)


@app.post("/upload")
async def upload_file(
    file: UploadFile = File(...),
    caption: str = Form(""),
    session: AsyncSession = Depends(get_async_session),
):
    temp_file_path = None
    try:
        with tempfile.NamedTemporaryFile(
            delete=False, suffix=os.path.splitext(file.filename)[1]
        ) as temp_file:
            temp_file_path = temp_file.name
            shutil.copyfileobj(file.file, temp_file)

        upload_result = imagekit.files.upload(
            file=temp_file_path,
            file_name=file.filename,
            tags=["backend_upload"],
            # file=open(temp_file_path, "rb"),
            # file_name=file.filename,
            # options=UploadFileRequestOptions(
            #     use_unique_file_name=True, tags=["backend_upload"]
            # ),
        )

        if upload_result.response_metadata.http_status_code == 200:
            post = Post(
                caption=caption,
                url=upload_result.url,
                file_type="video"
                if file.content_type.startswith("video/")
                else "image",
                file_name=upload_result.file_name,
            )
            session.add(post)  # add the post to the database
            await session.commit()  # save the changes to the database
            await session.refresh(post)  # refresh the post object
            return post
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    finally:
        if temp_file_path and os.path.exists(temp_file_path):
            os.unlink(temp_file_path)
        file.file.close()


@app.get("/feeds")
async def read_feeds(session: AsyncSession = Depends(get_async_session)):
    result = await session.execute(select(Post).order_by(Post.created_at.desc()))
    posts = [post[0] for post in result.all()]

    posts_data = []
    for post in posts:
        posts_data.append(
            {
                "id": post.id,
                "caption": post.caption,
                "url": post.url,
                "file_type": post.file_type,
                "file_name": post.file_name,
                "created_at": post.created_at.isoformat(),
            }
        )
    return {"posts": posts_data}
