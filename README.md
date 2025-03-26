# Line Notify Adapter

Line Notify Adapter is an adapter tool designed to migrate from LINE Notify to LINE Official Account (OA). It provides a seamless way to send notifications from your system to users via LINE OA groups, without disrupting your existing workflow.

## Features

- Migrate notifications from LINE Notify to LINE Official Account (OA)
- Supports sending notifications to LINE OA groups and other platforms
- Simple integration using HTTP API

## Prepare line OA account

1. Create line OA account
2. Enable line message api
3. Issue channel access token
4. Go to `Webhook.site` then copy url to setting Webhook URL (For debug group id or user id)
5. Set allow account to join groups and multi-person chats (For auto join group)
6. Invite OA bot to group
7. Get group id on `Webhook.site`
8. Disable Use webhook

## Installation

1. Clone the repository:

```bash
git clone https://github.com/larb26656/line-notify-adapter.git
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the service:

```bash
go run main.go
```

4. The service will be available on http://localhost:8080.

## Usage

### Authorization Token Format

The authorization token follows a specific pattern:

```
{{channelAccessToken}}_targetKey_{{tagetToken}}
```

- `{{channelAccessToken}}`: Your LINE Official Account access token
- `_targetKey_`: A separator key
- `{{tagetToken}}`: The target group or user ID

**Example Token:**

```
a1b2c3d4e5f6g7h8_TARGET_U1234567890
```

### Sending a Notification

To send a notification, use the following curl command:

```bash
curl --location 'http://localhost:8080/api/v1/notify' \
--header 'Authorization: Bearer a1b2c3d4e5f6g7h8_TARGET_U1234567890' \
--form 'message="test message"'
```

### Parameters

- **Authorization**: Follows the `Bearer {{channelAccessToken}}_targetKey_{{tagetToken}}` pattern
- **message**: The message to be sent to the user

## Notes

- At the moment, the API only supports sending plain text messages. More advanced notification features will be added in future releases.
