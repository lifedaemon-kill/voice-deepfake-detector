from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import uvicorn

from starlette.responses import JSONResponse
from transformers import pipeline

# Загрузка модели мелоди в пайплайн
predict1 = pipeline("audio-classification", model="pretrained_models/MelodyMachine")

predict2 = pipeline("audio-classification", model="mo-thecreator/Deepfake-audio-detection")

app = FastAPI()


class AudioPathRequest(BaseModel):
    file_path: str


@app.post("/predict")
async def predict(request: AudioPathRequest):
    """
    Ожидается json {file_path: string}
    РАЗДЕЛИТЕЛИ UNIX-LIKE path/to/file/file.ogg
    Возвращается вероятность того, что аудио является fake json {score1: float, score2: float, text: string}
    """
    file_path = request.file_path
    print(file_path)
    try:
        res1 = predict1(file_path)
        res2 = predict2(file_path)

    except FileNotFoundError as fnf_error:
        raise HTTPException(status_code=404, detail=str(fnf_error))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

    if res1[0]["label"] == "fake":
        score1 = res1[0]["score"]
    else:
        score1 = res1[1]["score"]

    if res2[0]["label"] == "fake":
        score2 = res2[0]["score"]
    else:
        score2 = res2[1]["score"]

    result = {
        "MelodyMachine": score1,
        "mo-thecreator": score2,
    }

    return JSONResponse(content=result)


if __name__ == "__main__":
    uvicorn.run("main:app", host="127.0.0.1", port=8090, reload=True)
