from fastapi import status, APIRouter
from .greeting_service import (
	greet
)

router = APIRouter(
	prefix="/greeting",
	tags=["greeting"],
	responses={404: {"description": "Not found"}}
)


@router.get(
	path="",
	response_description="Get a greeting",
	status_code=status.HTTP_200_OK)
async def greeting():
	return greet()
