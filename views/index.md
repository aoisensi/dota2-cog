Dota2 Cog
=========

Discordの役職に自動でDota2のランク(メダル)を割り当てるbotです

1時間に1回、自動でOpenDotaから情報を取ってきて役職を割り当てます

* * *

### 使い方

#### サーバーアドミン

[ここ](/add-bot)からサーバーにbotを追加してください

追加されたロールは色や名前の変更など自由に編集可能です(多分…)

#### サーバーメンバー

Dota2 Cog BotにDMで `/dota2-cog register` コマンドを送って

表示されたリンクをクリックし、Steamとの連携を有効にしてください

* * *

### コマンドリスト

##### /dota2-cog register

SteamとDiscordのIDを連携させます (DMのみ)

一度実行すれば、他のサーバーでこのbotを導入したとしても再使用は不要です

##### /dota2-cog force-fetch

現在のサーバーの全ユーザーのランクを取得します (乱用禁止!) (サーバーコマンドのみ) (管理者のみ)

##### /dota2-cog fix-rank-roles

ランクロールの修正をします 間違えてロールを削除してしまったときなどに使ってください (サーバーコマンドのみ) (管理者のみ)

##### /dota2-cog fix-registerd-role

登録済みロールの修正をします 間違えてロールを削除してしまったときなどに使ってください (サーバーコマンドのみ) (管理者のみ)

* * *

### バージョン履歴

*   #### 0.2.1 2020-06-21

    サーバーオーナー(冠が付いてるやつ)でも管理者コマンドが使えるようになった

    `register`をやっぱりDMのみにした(セキュリティ上の問題で)

    このページをMarkdownで書くようにした

*   #### 0.2.0 2020-06-21
    
    登録済みユーザーを割り当てる `Dota2 Cog Registerd` ロールを作った
    
    `/dota2-cog fix-roles` コマンドを `/dota2-cog fix-rank-roles` に変更
    
*   #### 0.1.4 2020-06-15
    
    force-fetchしたときにプログレスバーを表示するようにした
    
    fetch中にforce-fetchした場合エラーを返すようにした
    
    fix-roles コマンドを追加した
    
*   #### 0.1.3 2020-06-14
    
    ロールの再割り当てのアルゴリズムの最適化
    
*   #### 0.1.2 2020-06-13
    
    ちょっとだけ修正
    
*   #### 0.1.1 2020-06-12
    
    BotアカウントのFetchをスキップするようにした
    
*   #### 0.1.0 2020-06-10
    
    初期バージョン

* * *

### 寄付のお願い

Dota2 Cog Bot は無料で使えますが、もしあなたがこのbotの維持に協力してくれるなら寄付をお願いします

このサービスの運営には毎月のサーバー代などが掛かっています

#### 月額支援

[![Donate using Liberapay](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/aoisensi/donate)

#### 1回限りの寄付

<style>.bmc-button img{height: 34px !important;width: 35px !important;margin-bottom: 1px !important;box-shadow: none !important;border: none !important;vertical-align: middle !important;}.bmc-button{padding: 7px 15px 7px 10px !important;line-height: 35px !important;height:51px !important;text-decoration: none !important;display:inline-flex !important;color:#ffffff !important;background-color:#5F7FFF !important;border-radius: 5px !important;border: 1px solid transparent !important;padding: 7px 15px 7px 10px !important;font-size: 28px !important;letter-spacing:0.6px !important;box-shadow: 0px 1px 2px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 1px 2px 2px rgba(190, 190, 190, 0.5) !important;margin: 0 auto !important;font-family:'Cookie', cursive !important;-webkit-box-sizing: border-box !important;box-sizing: border-box !important;}.bmc-button:hover, .bmc-button:active, .bmc-button:focus {-webkit-box-shadow: 0px 1px 2px 2px rgba(190, 190, 190, 0.5) !important;text-decoration: none !important;box-shadow: 0px 1px 2px 2px rgba(190, 190, 190, 0.5) !important;opacity: 0.85 !important;color:#ffffff !important;}</style><link href="https://fonts.googleapis.com/css?family=Cookie" rel="stylesheet"><a class="bmc-button" target="_blank" href="https://www.buymeacoffee.com/aoisensi"><img src="https://cdn.buymeacoffee.com/buttons/bmc-new-btn-logo.svg" alt="Buy me a soda"><span style="margin-left:5px;font-size:28px !important;">Buy me a soda</span></a>

* * *

### その他

本当はDiscordのプロフィールからSteamIDを取得したかったがAPIがサポートしてないのでログインさせるという形になった…

1000人以上いるサーバーでは最初の1000人しかデータを取れない 後々何とかする

サーバーにいる人全員に連携を促すDMを送るコマンド作りたいがスパム扱いされそう…

ソースコード:[https://github.com/aoisensi/dota2-cog](https://github.com/aoisensi/dota2-cog)

なにか問題点などあれば**aoisensi#0634**または[https://steamcommunity.com/id/aoisensi](https://steamcommunity.com/id/aoisensi)までどうぞ

I will write this page with English.
