---
title: 待测试
description: 当前版本已实现但仍需人工验证的变更项
---

# 待测试

## 图片生成 SSE 响应兼容

- 已支持后端持久化图片任务从 `text/event-stream` 响应的 `data:` 事件中读取 `url` 或 `b64_json`，包括 `image_generation.partial_succeeded` 事件。
- 已支持前端直连图片接口从携带图片字段的 SSE 事件中提取结果，并按 `image_index` 保留每张图片的最新事件。
- 待人工确认：为 console 渠道开启图片“流式传输”后，生图工作台和无限画布均能正常显示生成图片，失败详情中不再出现 `invalid character 'e' looking for beginning of value`。

## 视频生成接口统一为 console 契约

- 无限画布视频创建与查询接口已统一为 `POST /v1/video/generations` 和 `GET /v1/video/generations/{taskId}`。
- 默认视频请求已改为 console 使用的 JSON 结构，通用参数位于顶层，比例、分辨率、声音、水印和参考素材位于 `metadata` / `metadata.content`。
- Seedance 模型不再仅凭模型名被误判为火山 Ark 直连；只有渠道地址明确属于 Ark/火山时才转换为 `/contents/generations/tasks`。
- KIE、APIMart、Ark 和 Agnes 继续通过渠道适配映射到各自原生路径。
- 已将上游 `NOT_START` / `NOT_STARTED` 状态归一化为排队中，并让后台继续轮询已保存为这些状态的任务，避免首次查询后永久停在 0%。
- 已支持从 console 视频任务响应的 `data.result_url` 和嵌套 `content.video_url` 读取最终视频地址，避免任务完成后被客户端误判为无结果。
- 待人工确认：console 渠道下 Seedance 文生视频、首帧/尾帧及多模态参考任务可以创建并持续轮询，不再返回 404 或 415。
