# Lumen

Lumen, **YAML tabanlı kurallar** ile çalışan, container içinde kolayca çalıştırılabilen hafif bir kural/analiz motorudur. Temel amacı, kullanıcı tarafından tanımlanan `rules.yaml` dosyasını okuyarak bu kurallara göre belirli kontrolleri, analizleri veya işlemleri gerçekleştirmektir.

Proje, **Docker** ile çalışacak şekilde tasarlanmıştır. Böylece herhangi bir ek bağımlılık kurmadan, sadece Docker kullanarak hızlıca ayağa kaldırılabilir.

---

## 🚀 Ne İşe Yarar?

* YAML formatında tanımlanmış kuralları okur
* Kurallara göre analiz veya kontrol işlemleri yapar
* Taşınabilir ve izole bir ortamda çalışır (Docker sayesinde)
* CI/CD, güvenlik kontrolleri veya otomatik denetimler için uygundur

> Projenin temel felsefesi: **"Kuralı dosyada tanımla, ortamdan bağımsız çalıştır."**

---

## 📦 Gereksinimler

* Docker (20.x veya üzeri önerilir)

Docker dışında sisteminizde herhangi bir ek kurulum gerekmez.

---

## ⚙️ Kurulum

### 1️⃣ Repoyu Klonlayın

```bash
git clone https://github.com/tuncaycelikkanat/lumen.git
cd lumen
```

---

### 2️⃣ Docker Image Oluşturma

Projeyi Dockerize etmek için aşağıdaki komutu kullanabilirsiniz:

```bash
docker build -t lumen:latest .
```

---

### 3️⃣ Çalıştırma

Lumen, çalışırken dışarıdan bir `rules.yaml` dosyasına ihtiyaç duyar. Bu dosya container içerisine volume olarak bağlanır.

Örnek çalıştırma komutu:

```bash
docker run --rm -it \
  -v $(pwd)/rules.yaml:/app/rules.yaml \
  lumen:latest
```

#### 🔍 Komut Açıklaması

* `--rm` : Container kapandığında otomatik silinir
* `-it` : Etkileşimli terminal
* `-v $(pwd)/rules.yaml:/app/rules.yaml` : Yerel `rules.yaml` dosyasını container içine bağlar
* `lumen:latest` : Kullanılacak Docker image

---

## 📝 rules.yaml Örneği

```yaml
rules:
  - name: Example Rule
    description: Örnek bir kural
    condition: value > 10
    action: warn
```

> ⚠️ `rules.yaml` içeriği ve desteklenen alanlar projenin gelişimine göre değişebilir.

---

## 🛠️ Geliştirme

Bu proje Go ile geliştirilmiştir ve Docker odaklı bir çalışma modeli benimser.

- Uygulamanın giriş noktası `cmd/` dizini altındadır
- Derleme sırasında statik bir binary (`lumen`) üretilir
- Çalışma zamanı yapılandırması `rules.yaml` dosyası üzerinden yapılır
- Docker, önerilen ve desteklenen çalışma ortamıdır

---

## 📌 Yol Haritası (Planlanan)

* [ ] Daha detaylı kural şeması
* [ ] JSON/YAML validation
* [ ] CLI parametreleri
* [ ] Çıktı formatları (JSON, text, report)
* [ ] CI/CD entegrasyon örnekleri

---

## 🤝 Katkı

Katkılar memnuniyetle karşılanır 🙌

1. Fork'layın
2. Yeni bir branch açın (`feature/my-feature`)
3. Değişikliklerinizi commit edin
4. Pull Request oluşturun

---

## 📄 Lisans

Bu proje MIT lisansı ile lisanslanmıştır.

---

## ✨ İletişim

Her türlü öneri ve geri bildirim için GitHub Issues üzerinden iletişime geçebilirsiniz.
