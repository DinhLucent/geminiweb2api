<p align="center">
  <img src="assets/gemini.png" width="400" alt="Gemini Logo">
</p>

<p align="center">
  <a href="https://github.com/DinhLucent/geminiweb2api/releases"><img src="https://img.shields.io/github/v/release/DinhLucent/geminiweb2api?style=flat-square&logo=github&color=3670ad" alt="Release"></a>
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="https://www.docker.com/"><img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker" alt="Docker"></a>
  <a href="https://github.com/DinhLucent/geminiweb2api/blob/main/LICENSE"><img src="https://img.shields.io/github/license/DinhLucent/geminiweb2api?style=flat-square&color=orange" alt="License"></a>
  <img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg?style=flat-square" alt="Maintained">
</p>

<p align="center">
  <a href="https://github.com/DinhLucent/geminiweb2api/stargazers"><img src="https://img.shields.io/github/stars/DinhLucent/geminiweb2api?style=flat-square&color=gold&label=stars" alt="Stars"></a>
  <a href="https://github.com/DinhLucent/geminiweb2api/issues"><img src="https://img.shields.io/github/issues/DinhLucent/geminiweb2api?style=flat-square&color=red&label=issues" alt="Issues"></a>
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome">
</p>

<p align="center">
  <a href="README.md">🇻🇳 Tiếng Việt</a> | <a href="README_en.md">🇬🇧 English</a>
</p>

<h1 align="center">Gemini Web To API</h1>

<p align="center">
  Chuyển đổi giao diện web Google Gemini thành REST API tiêu chuẩn.<br/>
  Sử dụng khả năng xử lý của Gemini thông qua cookie trình duyệt, không cần API key.
</p>

<p align="center">
  <b>Tác giả: <a href="https://github.com/DinhLucent">DinhLucent (Eng.Dinh)</a></b><br/>
  <sub>Nguyên tác: <a href="https://github.com/ntthanh2603">@ntthanh2603</a></sub>
</p>

> [!NOTE]
> Dự án này phục vụ mục đích nghiên cứu và học tập. Vui lòng không sử dụng cho mục đích thương mại.

> [!WARNING]
> Dự án không liên kết với Google. Hệ thống sử dụng kỹ thuật đảo ngược cookie web và có thể không tuân thủ [Điều khoản dịch vụ của Google](https://policies.google.com/terms). Người dùng tự chịu trách nhiệm về các rủi ro liên quan đến tài khoản hoặc dữ liệu khi vận hành.

---

## Tổng quan

Gemini Web To API khởi tạo máy chủ proxy cục bộ kết nối trực tiếp đến giao diện web Google Gemini bằng cookie phiên đăng nhập từ trình duyệt. Máy chủ cung cấp các endpoint tương thích với chuẩn gọi API của OpenAI, Claude và Gemini.

Dự án phù hợp cho các nhu cầu:
- Tích hợp mô hình Gemini vào các ứng dụng hiện có mà không cần API key trả phí.
- Thử nghiệm cục bộ các tính năng mới nhất trên web Gemini.
- Sử dụng trực tiếp các SDK phổ biến của OpenAI, LangChain hoặc Google GenAI.

---

## Hướng dẫn cài đặt

### Cách 1: Chạy bằng Docker (Khuyên dùng)

Cách này không cần tải mã nguồn hay cài đặt môi trường Go, hình ảnh container được tải trực tiếp từ GitHub Container Registry.

#### Bước 1: Lấy cookie tài khoản

> [!WARNING]
> Tuyệt đối không chia sẻ hoặc đẩy chuỗi cookie lên kho lưu trữ công khai vì giá trị này cho phép truy cập trực tiếp vào tài khoản Google của bạn.

1. Đăng nhập vào [gemini.google.com](https://gemini.google.com) trên trình duyệt.
2. Nhấn `F12` để mở Công cụ cho nhà phát triển, chuyển sang thẻ **Network**, tải lại trang và chọn một yêu cầu có tên `batchexecute?rpcids=otAQ7b`.
3. Sao chép toàn bộ giá trị trong mục Request Header `Cookie`. Nếu đường dẫn yêu cầu có dạng `/u/2/`, ghi lại số `2` để gán cho biến `GEMINI_AUTH_USER`.

#### Bước 2: Khởi chạy container

```bash
docker run -d -p 4981:4981 \
  -e GEMINI_COOKIES="your_complete_cookie_header" \
  -e GEMINI_AUTH_USER="2" \
  -e GEMINI_REFRESH_INTERVAL=1440 \
  -e GEMINI_MAX_RETRIES=3 \
  -e GEMINI_TEMPORARY=false \
  -e APP_ENV=production \
  -e RATE_LIMIT_ENABLED=true \
  -e RATE_LIMIT_WINDOW_MS=60000 \
  -e RATE_LIMIT_MAX_REQUESTS=10 \
  -v ./cookies:/home/appuser/.cookies \
  --tmpfs /tmp:rw,size=512m \
  --tmpfs /home/appuser/.cache:rw,size=256m \
  --name gemini-web-to-api \
  --restart unless-stopped \
  ghcr.io/ntthanh2603/gemini-web-to-api:latest
```

Sau khi chạy lệnh, chuyển tiếp đến phần [Kiểm tra hoạt động](#kiểm-tra-hoạt-động).

---

### Cách 2: Chạy từ mã nguồn

Phù hợp khi bạn cần chỉnh sửa mã nguồn hoặc biên dịch cho các kiến trúc phần cứng đặc thù.

#### Bước 1: Tải mã nguồn

```bash
git clone https://github.com/DinhLucent/geminiweb2api.git
cd geminiweb2api
```

#### Bước 2: Cấu hình tệp môi trường

> [!WARNING]
> Không đưa tệp `.env` vào git commit để tránh lộ thông tin xác thực tài khoản.

1. Đăng nhập [gemini.google.com](https://gemini.google.com).
2. Nhấn `F12`, vào thẻ **Network**, tải lại trang và tìm yêu cầu `batchexecute?rpcids=otAQ7b`.
3. Sao chép toàn bộ giá trị header `Cookie` vào biến `GEMINI_COOKIES`.
4. Nếu URL có dạng `/u/<số>/`, điền số đó vào `GEMINI_AUTH_USER`. Nếu không có, để trống biến này. Luôn lấy cookie và chỉ số tài khoản từ cùng một tab trình duyệt.
5. Tạo tệp `.env` từ tệp mẫu:

   ```bash
   cp .env.example .env
   ```

6. Cập nhật các thông số vào `.env`:

   ```env
   GEMINI_COOKIES=your_complete_cookie_header
   GEMINI_AUTH_USER=2
   GEMINI_REFRESH_INTERVAL=1440
   GEMINI_MAX_RETRIES=3
   GEMINI_TEMPORARY=false
   APP_ENV=production
   RATE_LIMIT_ENABLED=true
   RATE_LIMIT_WINDOW_MS=60000
   RATE_LIMIT_MAX_REQUESTS=10
   ```

#### Bước 3: Khởi động dịch vụ

| Phương pháp | Lệnh thực thi | Yêu cầu hệ thống |
| --- | --- | --- |
| Docker Compose | `docker compose up -d --build` | Docker |
| Go trực tiếp | `go run cmd/server/main.go` | [Go 1.21+](https://golang.org/dl/) |
| Task (chế độ phát triển) | `task dev` | [Task](https://taskfile.dev) và [Air](https://github.com/air-verse/air) (`task install-air`) |

Lệnh `task dev` kích hoạt chế độ tự động tải lại (hot reload) khi có thay đổi trong các tệp `.go`, `.json` hoặc `.env`.

---

### Kiểm tra hoạt động

Gửi thử một yêu cầu qua cURL:

```bash
curl -X POST http://localhost:4981/openai/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model": "gemini-advanced", "messages": [{"role": "user", "content": "Xin chào!"}]}'
```

Dịch vụ mặc định lắng nghe tại địa chỉ `http://localhost:4981`.

---

## Tính năng

- Đa giao thức: Hỗ trợ đồng thời ba giao thức OpenAI (`/openai/v1`), Claude (`/claude`) và Gemini (`/gemini/v1beta`).
- Tương thích SDK: Dùng trực tiếp với thư viện chính thức của OpenAI, Claude và Google GenAI mà không cần cấu hình phức tạp.
- Quản lý phiên làm việc: Tự động xoay vòng cookie và duy trì token phiên để giữ kết nối ổn định.
- Hiệu năng hệ thống: Viết bằng ngôn ngữ Go trên nền Fiber.
- Sẵn sàng triển khai: Hỗ trợ Docker Compose, kiểm tra tình trạng dịch vụ (health check) và tài liệu API tương tác tại `/docs`.
- Tạo hình ảnh: Tương thích chuẩn tạo ảnh của OpenAI, trả về định dạng `b64_json` có xác thực.
- Xử lý ảnh đầu vào: Hỗ trợ truyền ảnh qua liên kết URL từ xa, chuỗi `data:` base64 và hỗ trợ nhiều ảnh tham chiếu cùng lúc.
- Tạo video: Sinh video định dạng MP4 (tỷ lệ 16:9 hoặc 9:16) thông qua công cụ Videos của Gemini Web, tương thích với SDK của OpenAI (`client.videos`) và Google GenAI (`generate_videos`).
- Chế độ Canvas: Thao tác với tài liệu và tệp mã nguồn Canvas qua các mã mô hình có hậu tố `-canvas` (ví dụ `gemini-pro-canvas`).
- Tạo âm thanh: Tích hợp công cụ âm thanh qua giao thức `client.audio.speech` của OpenAI và các mô hình `-music` của Gemini.

Xem thêm tài liệu chi tiết cho từng tính năng:
- [Tạo ảnh và xử lý ảnh đầu vào](docs/image-generation.md)
- [Tạo video](docs/video-generation.md)
- [Chế độ Canvas](docs/canvas.md)
- [Tạo âm thanh](docs/music.md)

---

## Cấu hình hệ thống

### Danh sách biến môi trường

| Tên biến | Bắt buộc | Mặc định | Ý nghĩa |
| --- | --- | --- | --- |
| `GEMINI_COOKIES` | Có | — | Toàn bộ header Cookie lấy từ tab Gemini Web |
| `GEMINI_AUTH_USER` | Không | — | Vị trí tài khoản Google trong URL (ví dụ `2` nếu URL là `/u/2/`) |
| `GEMINI_REFRESH_INTERVAL` | Không | `1440` | Chu kỳ xoay vòng cookie tính theo phút (tối thiểu `10`) |
| `GEMINI_MAX_RETRIES` | Không | `3` | Số lần thử lại tối đa khi yêu cầu thất bại |
| `GEMINI_TEMPORARY` | Không | `false` | Bật chế độ ẩn danh (không lưu lịch sử) cho mọi yêu cầu |
| `PORT` | Không | `4981` | Cổng dịch vụ lắng nghe |
| `RATE_LIMIT_ENABLED` | Không | `false` | Bật hoặc tắt tính năng giới hạn tần suất yêu cầu |
| `RATE_LIMIT_WINDOW_MS` | Không | `60000` | Khung thời gian giới hạn tần suất tính bằng mili-giây |
| `RATE_LIMIT_MAX_REQUESTS` | Không | `10` | Số yêu cầu tối đa cho phép trong mỗi khung thời gian |

### Hạn chế bị chặn yêu cầu (`GEMINI_REFRESH_INTERVAL`)

Máy chủ định kỳ xoay vòng cookie `__Secure-1PSIDTS` và làm mới token phiên của Gemini để duy trì kết nối. Việc xoay vòng quá thường xuyên có thể bị Google nhận diện là lưu lượng truy cập bất thường, dẫn đến lỗi 429 hoặc yêu cầu đăng nhập lại tài khoản.

- Giá trị mặc định là 1440 phút (24 giờ). Các giá trị dưới 10 phút sẽ được tự động điều chỉnh về 10. Giá trị bằng 0 hoặc số âm sẽ quay về mặc định.
- Nếu gặp lỗi giới hạn tần suất, giữ nguyên giá trị mặc định và bật `RATE_LIMIT_ENABLED=true` kèm giới hạn `RATE_LIMIT_MAX_REQUESTS` phù hợp.
- Nếu phiên làm việc hết hạn trước thời gian xoay vòng (nhật ký ghi nhận `marking client unhealthy`), hãy sao chép lại chuỗi cookie mới từ trình duyệt.

### Thứ tự ưu tiên cấu hình

1. Biến môi trường hệ thống (ưu tiên cao nhất)
2. Tệp `.env`
3. Giá trị mặc định trong mã nguồn (ưu tiên thấp nhất)

---

## Hướng dẫn sử dụng

### Cơ chế phát hiện mô hình

Dịch vụ tự động đồng bộ danh sách mô hình từ tài khoản đăng nhập khi khởi tạo và làm mới phiên. Kiểm tra các mô hình khả dụng bằng lệnh:

```bash
curl http://localhost:4981/openai/v1/models
```

Các mã mô hình phản ánh đúng nhãn hiển thị hiện thời trên giao diện web (ví dụ `gemini-3.6-flash`, `gemini-3.1-pro`) thay vì sử dụng danh sách cố định trong mã nguồn. Khi Google cập nhật hoặc đổi tên mô hình, các mô hình mới sẽ tự động xuất hiện. Tên `gemini-advanced` được duy trì như một bí danh tương thích cho dòng Pro.

Trường `model` trong phản hồi trả về phiên bản cụ thể đang xử lý yêu cầu. Trong trường hợp tài khoản hết hạn mức tạm thời của mô hình được chọn, web Gemini có thể chuyển sang phục vụ bằng mô hình cấp thấp hơn. Proxy sẽ nhận biết và hiển thị đúng mô hình thực tế trong trường `model`, đồng thời lưu mã mô hình được yêu cầu ban đầu tại trường `requested_model`.

### Sử dụng với OpenAI SDK (Python)

```python
from openai import OpenAI

client = OpenAI(
    base_url="http://localhost:4981/openai/v1",
    api_key="not-needed"
)

response = client.chat.completions.create(
    model="gemini-advanced",
    messages=[{"role": "user", "content": "Xin chào!"}]
)
print(response.choices[0].message.content)
```

### Sử dụng với Claude SDK (LangChain Python)

```python
from langchain_anthropic import ChatAnthropic

llm = ChatAnthropic(
    base_url="http://localhost:4981/claude",
    model="gemini-advanced",
    api_key="not-needed"
)

response = llm.invoke("Giải thích ngắn gọn về điện toán lượng tử")
print(response.content)
```

### Sử dụng với Google GenAI SDK (Python)

```python
from google import genai

client = genai.Client(
    api_key="not-needed",
    http_options={
        "base_url": "http://localhost:4981/gemini",
        "api_version": "v1beta",
    },
)

response = client.models.generate_content(
    model="gemini-advanced",
    contents="Viết một đoạn thơ ngắn về lập trình",
)
print(response.text)
```

Cài đặt thư viện bằng lệnh `pip install google-genai`. Dịch vụ hỗ trợ cả gọi thông thường và luồng phản hồi liên tục qua `generate_content_stream`.

### Gọi trực tiếp bằng cURL

```bash
curl -X POST http://localhost:4981/openai/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-advanced",
    "messages": [{"role": "user", "content": "Trí tuệ nhân tạo là gì?"}],
    "stream": false
  }'
```

Xem thêm các ví dụ khác trong thư mục [`examples/`](examples/).

---

## Tài liệu API

Sau khi khởi động dịch vụ, truy cập `http://localhost:4981/docs` để xem và tương tác với tài liệu API qua giao diện Scalar.

![Scalar UI](assets/swagger.png)

---

## Đóng góp mã nguồn

Mọi đóng góp nhằm hoàn thiện dự án đều được hoan nghênh qua Pull Request.

1. Fork kho lưu trữ về tài khoản cá nhân
2. Tạo nhánh tính năng mới (`git checkout -b feature/ten-tinh-nang`)
3. Lưu các thay đổi (`git commit -m 'Them tinh nang'`)
4. Đẩy nhánh lên GitHub (`git push origin feature/ten-tinh-nang`)
5. Tạo Pull Request vào nhánh chính

---

## Tác giả & Đội ngũ

- **Tác giả (Custom & Maintainer):** [DinhLucent](https://github.com/DinhLucent) (Eng.Dinh)
- **Nguyên tác (Original Author):** Nguyễn Tuấn Thành ([@ntthanh2603](https://github.com/ntthanh2603))

---

## Giấy phép

Dự án được phân phối theo giấy phép MIT. Xem tệp [LICENSE](LICENSE) để biết thêm chi tiết.
