# onlineclassbot
🏡 온라인 클래스용 봇
![커버 사진](https://raw.githubusercontent.com/cjaewon/onlineclassbot/main/media/cover.png)

## Features
- 시간표 알림 기능

## How to Use
```sh
go run main.go
```
### 설정파일
`.onlineclassbot.toml` 파일을 통해서 봇을 설정 할 수 있습니다.
```toml
[schedules] # 시간표
  [schedules."35 8 * * 1-5"]
    teacher = "담임 ***"
    url = "https://us04web.zoom.us/j/abc123"

  [schedules."20 16 * * 1-5"]
    teacher = "담임 ***"
    url = "https://us04web.zoom.us/j/abc123"

  # ..

[bot]
  token = "" # 디스코드 봇 토큰
  notify_id = "" # 시간표를 보낼 채널 id
```