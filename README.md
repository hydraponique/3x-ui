<div align="center">

<table width="100%">
<tr>
<td align="center" width="25%"><a href="https://github.com/hydraponique/roscomvpn-geoip">RoscomVPN GeoIP</a></td>
<td align="center" width="25%"><a href="https://github.com/hydraponique/roscomvpn-geosite">RoscomVPN Geosite</a></td>
<td align="center" width="25%"><a href="https://github.com/hydraponique/3x-ui"><b>🚀 3x-ui · RoscomVPN</b></a></td>
<td align="center" width="25%"><a href="https://github.com/hydraponique/roscomvpn-routing">RoscomVPN Routing</a></td>
</tr>
<tr>
<td align="center"><img src="https://img.shields.io/github/downloads/hydraponique/roscomvpn-geoip/total.svg" alt="Downloads"> <img src="https://data.jsdelivr.com/v1/package/gh/hydraponique/roscomvpn-geoip/badge" alt="jsDelivr"></td>
<td align="center"><img src="https://img.shields.io/github/downloads/hydraponique/roscomvpn-geosite/total.svg" alt="Downloads"> <img src="https://data.jsdelivr.com/v1/package/gh/hydraponique/roscomvpn-geosite/badge" alt="jsDelivr"></td>
<td align="center"><img src="https://img.shields.io/github/v/release/hydraponique/3x-ui.svg" alt="Release"> <img src="https://img.shields.io/github/downloads/hydraponique/3x-ui/total.svg" alt="Downloads"> <img src="https://img.shields.io/badge/Docker-ghcr.io-blue.svg?logo=docker" alt="Docker"></td>
<td align="center"><img src="https://img.shields.io/github/stars/hydraponique/roscomvpn-routing.svg" alt="Stars"> <img src="https://img.shields.io/badge/Happ-blue.svg" alt="Happ"> <img src="https://img.shields.io/badge/Mihomo-grey.svg" alt="Mihomo"> <img src="https://img.shields.io/badge/Incy-darkgreen.svg" alt="Incy"></td>
</tr>
</table>

# 🚀 3x-ui · RoscomVPN edition

### Форк [3x-ui](https://github.com/MHSanaei/3x-ui), в который встроен свежий [RoscomVPN-роутинг](https://github.com/hydraponique/roscomvpn-routing)

**Для кого:** хочется поднять свой VPN под 🇷🇺 / 🇧🇾 без танцев с правилами роутинга и обновлением geo-файлов.
**Сохраняет всё:** база, конфиги, клиенты, подписки, сертификаты — переходите на форк и обратно одной командой.

</div>

---

## 🎁 Что добавлено

<table width="100%">
<tr>
<td width="33%" align="center">

### 📡 Routing в подписку

В Happ-клиент сразу прилетает готовый профиль маршрутизации.

`DEFAULT` · `JSONSUB` · `WHITELIST` · своя ссылка

Меняется в один клик.

</td>
<td width="33%" align="center">

### 🌎 Geo-файлы из коробки

`geoip_ROSCOM.dat` и `geosite_ROSCOM.dat` уже в образе.

Кастомные RU/BY CIDR, "казённые" сервисы, Apple Push, минус РКН-листы.

</td>
<td width="33%" align="center">

### 🔄 Автообновление

Включил тумблер — раз в сутки сам тянет свежие geo-файлы и (если изменились) перезапускает Xray.

Никаких кронов в системе.

</td>
</tr>
</table>

### Ещё мелочи

- 🎯 В выпадашках Xray-правил появились пресеты `🚀 RoscomVPN ...` для IP, доменов, блок-листов и сервисов.
- 📊 На дашборде раздел **Geofiles** показывает размер и время последнего обновления каждого `.dat`.
- 🟢 На странице логина и в шапке дашборда — метка **🚀 RoscomVPN**, чтобы случайно не перепутать с апстримом.

---

## 🔍 Как понять, какая у вас сейчас версия стоит

Зайдите в SSH и наберите:

```bash
x-ui
```

<table width="100%">
<tr>
<td width="50%" align="center">

### ✅ Открылось меню

```
1. Install x-ui
2. Update x-ui
3. ...
```

➜ У вас **Standalone**

Обновляйтесь способом **1** или **2**

</td>
<td width="50%" align="center">

### ❌ `command not found`

```
bash: x-ui: command not found
```

➜ У вас **Docker** (или ничего)

Обновляйтесь способом **3**

</td>
</tr>
</table>

---

## 📦 Установка

> [!TIP]
> Всё, что у вас уже настроено — **остаётся**. База `/etc/x-ui/x-ui.db`, сертификаты, inbounds, клиенты, подписки — не трогаются.

<details>
<summary><h3>🟢 Способ 1 · Тихое обновление (для standalone — рекомендуемый)</h3></summary>

**Лучший выбор**, если у вас уже стоит апстрим или предыдущая версия форка. Не задаёт вопросов, ничего не сбрасывает — просто берёт логин/пароль/порт из существующей базы и обновляет панель.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/hydraponique/3x-ui/main/update.sh)
```

| Шаг | Действие |
|---|---|
| 1 | Остановит панель |
| 2 | Скачает свежий релиз форка |
| 3 | Распакует, поставит на место |
| 4 | Запустит панель обратно |

База, сертификаты, логин/пароль, порт — **всё остаётся**.

</details>

<details>
<summary><h3>🟡 Способ 2 · Интерактивная установка/обновление (для standalone)</h3></summary>

То же что способ 1, но **спросит** username / password / порт панели — удобно, если ставите впервые или хотите поменять учётные данные.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/hydraponique/3x-ui/main/install.sh)
```

База, сертификаты, inbounds, клиенты, подписки — **остаются** (как и в способе 1). Меняется только бинарь, geo-файлы и `x-ui.sh`. Если на сервере уже что-то стоит — это просто обновление с возможностью переустановить логин/пароль.

</details>

<details>
<summary><h3>🐳 Способ 3 · Docker / Compose</h3></summary>

Откройте ваш `docker-compose.yml` и поменяйте одну строчку:

```diff
services:
  3xui:
-   image: ghcr.io/mhsanaei/3x-ui:latest
+   image: ghcr.io/hydraponique/3x-ui:latest
```

Затем:

```bash
docker compose pull && docker compose up -d
```

Всё остальное (volumes, env, network) трогать не надо. Доступные теги: `latest`, `5.8.x`. Архитектуры: `amd64`, `arm64`, `armv7`, `armv6`, `386`.

</details>

---

## 🔁 Не понравилось — откатываемся

Standalone:
```bash
bash <(curl -Ls https://raw.githubusercontent.com/MHSanaei/3x-ui/main/install.sh)
```

Docker — поменять `image:` обратно на `ghcr.io/mhsanaei/3x-ui:latest` и `docker compose pull && up -d`.

База остаётся, ничего не теряется.

---

## ⚠️ Важно знать после установки

> [!WARNING]
> После обновления **Happ-клиенты сразу получат RoscomVPN-роутинг по умолчанию**. Это и есть смысл форка, но если хочется как раньше — зайдите в **Subscription Settings** → **Routing source** → выберите `Custom link` и оставьте поле пустым. ИЛИ выключите тумблер `Enable routing`. Применяется без перезапуска.

---

## 🙏 Спасибо

Этот форк — лишь надстройка. Вся базовая мощь панели — заслуга **[@MHSanaei](https://github.com/MHSanaei)** и его проекта **[3x-ui](https://github.com/MHSanaei/3x-ui)** (GPL-3.0).

Если RoscomVPN-надстройка вам не нужна — ставьте оригинал и поддержите автора:

<a href="https://www.buymeacoffee.com/MHSanaei" target="_blank">
  <img src="https://img.shields.io/badge/Buy%20MHSanaei-a%20coffee-yellow?style=for-the-badge&logo=buymeacoffee" alt="Buy MHSanaei a coffee">
</a>

#### Источники geo-данных

- [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat) — основной `geoip` / `geosite`
- [chocolate4u/Iran-v2ray-rules](https://github.com/chocolate4u/Iran-v2ray-rules) — `geoip_IR` / `geosite_IR`
- [runetfreedom/russia-v2ray-rules-dat](https://github.com/runetfreedom/russia-v2ray-rules-dat) — `geoip_RU` / `geosite_RU`
- [hydraponique/roscomvpn-geoip](https://github.com/hydraponique/roscomvpn-geoip) + [roscomvpn-geosite](https://github.com/hydraponique/roscomvpn-geosite) — `geoip_ROSCOM` / `geosite_ROSCOM`
- [hydraponique/roscomvpn-routing](https://github.com/hydraponique/roscomvpn-routing) — DEEPLINK-профили для Happ

---

<div align="center">

### Понравилось — ставь ⭐

**Сделано с ❤️ к свободному интернету**

`USDT TRC20: TMu3N2ZjK5omJ7n3WAj5MNCSM5querBXsR`

</div>
