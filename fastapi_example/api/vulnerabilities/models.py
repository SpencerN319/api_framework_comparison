from pydantic import (
	BaseModel,
	Field,
)
from uuid import uuid4


class Vulnerability(BaseModel):
	id: str = Field(default_factory=uuid4, alieas="_id")
	serverUrl: str
	taskId: str
	status: str
	properties: dict

	class Config:
		allow_population_by_field_name = True
		schema_extra = {
			"example": {
				"_id": "066de609-b04a-4b30-b46c-32537c7f1f6e",
				"serverUrl": "https://localhost:1234",
				"taskId": "1",
				"status": "SUCCESS",
				"properties": {}
			}
		}


class VulnerabilityUpdate(BaseModel):
	serverUrl: str
	taskId: str
	status: str
	properties: dict

	class Config:
		allow_population_by_field_name = True
		schema_extra = {
			"example": {
				"serverUrl": "https://localhost:1234",
				"taskId": "1",
				"status": "SUCCESS",
				"properties": {}
			}
		}
