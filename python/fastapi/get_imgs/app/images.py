import os

from dotenv import load_dotenv
from imagekitio import ImageKit

load_dotenv()

imagekit = ImageKit(
    private_key=os.environ["IMAGEKIT_PRIVATE_KEY"],
    # os.getenv("IMAGEKIT_PRIVATE_KEY"),
    # os.getenv("IMAGEKIT_PUBLIC_KEY"),
    # os.getenv("IMAGEKIT_URL"),
)
IMAGEKIT_PUBLIC_KEY = os.environ["IMAGEKIT_PUBLIC_KEY"]
IMAGEKIT_URL = os.environ["IMAGEKIT_URL"]
