package bot

const licence = `
Модели распознавания спуфинга аудио:

м1: MelodyMachine/Deepfake-audio-detection-V2
Заявленная точность: 0.9973
Лицензия: Apache License 2.0
Источник: https://huggingface.co/MelodyMachine/Deepfake-audio-detection-V2
-----------------------------------------
м2: mo-thecreator/Deepfake-audio-detection
Заявленная точность: 0.9882
Лицензия: Apache License 2.0
Источник: https://huggingface.co/mo-thecreator/Deepfake-audio-detection
-----------------------------------------
Apache License 2.0:
Лицензия открытого исходного кода, даёт пользователю право использовать программное обеспечение для любых целей, свободно изменять и распространять изменённые копии, за исключением названия
`

const help = `
Бот специализируется на распознавании подделки голоса (voice spoofing / deepfake)
Поддерживаются:
- Голосовые сообщения
- Файлы .ogg, .mp3

Команды:
/help - Показать это сообщение
/licence - Показать лицензии используемого ПО
`
