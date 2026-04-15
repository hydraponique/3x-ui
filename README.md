<div align="center">

<table width="100%">
<tr>
<td align="center"><a href="https://github.com/hydraponique/roscomvpn-geoip">RoscomVPN GeoIP</a></td>
<td align="center"><a href="https://github.com/hydraponique/roscomvpn-geosite">RoscomVPN Geosite</a></td>
<td align="center"><a href="https://github.com/hydraponique/roscomvpn-routing">RoscomVPN Routing</a></td>
<td align="center"><b>🚀 3x-ui · RoscomVPN edition</b></td>
</tr>
</table>

# 🚀 3x-ui · RoscomVPN edition

[![Release](https://img.shields.io/github/v/release/hydraponique/3x-ui.svg?style=flat-square)](https://github.com/hydraponique/3x-ui/releases)
[![Downloads](https://img.shields.io/github/downloads/hydraponique/3x-ui/total.svg?style=flat-square)](https://github.com/hydraponique/3x-ui/releases/latest)
[![Docker](https://img.shields.io/badge/Docker-ghcr.io%2Fhydraponique-blue?style=flat-square&logo=docker)](https://github.com/hydraponique/3x-ui/pkgs/container/3x-ui)
[![License](https://img.shields.io/badge/License-GPL%20v3-blue.svg?style=flat-square)](https://www.gnu.org/licenses/gpl-3.0.en.html)

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
<summary><h3>🟢 Способ 1 · Одна команда (для standalone)</h3></summary>

Самый простой. Подходит для свежей установки и для апгрейда с апстрима.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/hydraponique/3x-ui/main/install.sh)
```

Что сделает:

| Шаг | Действие |
|---|---|
| 1 | Остановит панель |
| 2 | Скачает свежий релиз форка |
| 3 | Распакует, поставит на место |
| 4 | Запустит панель обратно |

База и сертификаты остаются.

</details>

<details>
<summary><h3>🟡 Способ 2 · Тихое обновление (для standalone, без вопросов)</h3></summary>

То же что способ 1, но **не спрашивает** username/password/port — берёт из существующей базы. Удобно, если уже всё настроено и нужен просто апдейт.

```bash
bash <(curl -Ls https://raw.githubusercontent.com/hydraponique/3x-ui/main/update.sh)
```

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
