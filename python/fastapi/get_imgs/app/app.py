from fastapi import FastAPI, HTTPException

app = FastAPI()


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
def read_post(post_id: int):  # add path param
    if post_id not in posts:
        raise HTTPException(status_code=404, detail="Post not found")
    return posts.get(post_id)
