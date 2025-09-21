# 🌤️ Weather App

Một ứng dụng web đơn giản hiển thị thông tin thời tiết theo thành phố,  
được xây dựng bằng **vanilla Golang** cho backend API và **vanilla JavaScript** cho frontend.

---

## ✨ Tính năng
- Tra cứu thời tiết theo tên thành phố.
- Hiển thị nhiệt độ, điều kiện thời tiết, và icon minh họa.
- Backend viết bằng **Go** (không dùng framework).
- Frontend viết bằng **HTML + CSS + JS** (không dùng framework).
- Kết nối tới OpenWeatherMap API (hoặc API thời tiết khác).

---

## 🛠️ Công nghệ sử dụng
### Backend
- [Golang](https://go.dev/) (phiên bản >= 1.24.2)  
- HTTP server tích hợp sẵn (`net/http`).  
- Go module để quản lý dependencies.  

### Frontend
- HTML5, CSS3 cơ bản.  
- Vanilla JavaScript (`fetch` API để gọi backend).

### Triển khai
- Docker + Docker Compose để chạy toàn bộ app.  
- Alpine Linux base image cho image nhỏ gọn.  
