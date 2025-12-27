# defines the type of data, the various endpoints will accept

from pydantic import BaseModel


# use PostCreate to provide schema for the request body of create_post endpoint
class PostCreate(BaseModel):
    title: str
    caption: str


# use PostResponse to provide schema for the response of create_post endpoint
# helps with bett
class PostResponse(BaseModel):
    title: str
    caption: str
