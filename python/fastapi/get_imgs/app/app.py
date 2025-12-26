from contextlib import asynccontextmanager

from fastapi import FastAPI, HTTPException
from sqlalchemy.ext.asyncio import AsyncSession

from app.db import Post, create_db, get_async_session
from app.schemas import PostCreate, PostResponse


@asynccontextmanager
async def lifespan(app: FastAPI):
    await create_db()
    yield


# create the FastAPI app, and add the lifespan context manager
app = FastAPI(lifespan=lifespan)


posts = {
    1: {"title": "First Post", "body": "This is my first post!"},
    2: {"title": "Second Post", "body": "This is my second post!"},
    3: {"title": "Third Post", "body": "This is my third post!"},
    4: {"title": "Fourth Post", "body": "This is my fourth post!"},
    5: {"title": "Fifth Post", "body": "This is my fifth post!"},
    6: {"title": "Sixth Post", "body": "This is my sixth post!"},
    7: {"title": "Seventh Post", "body": "This is my seventh post!"},
    8: {"title": "Eighth Post", "body": "This is my eighth post!"},
}


@app.get("/posts")
def read_posts(limit: int = None):  # add query param
    # default value is None to make it optional
    if limit:
        return list(posts.values())[:limit]
    return posts


@app.get("/posts/{post_id}")
def read_post(post_id: int) -> PostResponse:  # add path param
    if post_id not in posts:
        raise HTTPException(status_code=404, detail="Post not found")
    return posts.get(post_id)


@app.post("/posts")
def create_post(post: PostCreate) -> PostResponse:  # request body
    # to accept the request body data, need a schema in fastapi
    new_post = {"title": post.title, "body": post.body}
    posts[max(posts.keys()) + 1] = new_post
    return new_post


@app.put("/posts/{post_id}")
def update_post(post_id: int, post: PostCreate) -> PostResponse:  # add path param
    if post_id not in posts:
        raise HTTPException(status_code=404, detail="Post not found")
    posts[post_id] = post


@app.delete("/posts/{post_id}")
def delete_post(post_id: int):  # add path param
    if post_id not in posts:
        raise HTTPException(status_code=404, detail="Post not found")
    del posts[post_id]
    return {"message": f"Post {post_id} deleted"}
