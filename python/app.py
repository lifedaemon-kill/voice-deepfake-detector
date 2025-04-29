from fastapi import FastAPI, Header, HTTPException
from fastapi.responses import JSONResponse

app = FastAPI()

@app.get("/predict")
async def predict(x_audio_path: str = Header(None)):
    """
    Эндпоинт /predict:
      - Ожидает наличие заголовка X_audio_path
      - Возвращает фиктивный JSON с предсказаниями:
          { "model1": <float>, "model2": <float> }
    """
    if x_audio_path is None:
        raise HTTPException(status_code=400, detail="Отсутствует обязательный заголовок: X-audio-path")

    # Выводим полученное значение заголовка (для отладки)
    print(f"Получен заголовок X-audio_path: {x_audio_path}")

    # TODO сдесь нужно сделать два запроса моделям
    # predict1 = pythorch speech brain
    # predict2 = ASVSpoof2021

    result = {
        "model1": 0.80,
        "model2": 0.90
    }
    return JSONResponse(content=result)


@app.get("/")
async def root():
    return {"message": "Сервер для предсказаний. Используйте эндпоинт /predict"}
