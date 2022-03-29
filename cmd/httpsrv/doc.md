# WebSocketのライフサイクル
- FEからwsを開けるHTTPリクエスト
- BEは`Connection: upgrade`のレスと、ws通信用のkeyを返す
  - この時点からfor loopとselectでchannelを使って、messageの送信と受信を始める
- 終わる時は、errorをdetectしたら
  - channelを閉じる
  - hubからclientを削除
    - このoperationはchannelを使って、hubまでmessageを送っている
