# CIOS OpenAPI Files

## How to Update

1. ./openapis/配下のOpenAPIを編集
2. main.pyを実行(sync openapi.yml)
```shell script
python main.py
```

## OpenAPI Yaml

| API | FILE | VERSION | Source |
|---|---|---|---|
|ALL API| [openapi.yml](openapis/openapi.yml) | --- |
|Account API| [account.openapi.yml](./openapis/account.openapi.yml) | 0.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/account/) |
|Collection API| [collection.openapi.yml](./openapis/collection.openapi.yml) | 0.0.0 | [portal](https://de9nof8bnxm56.cloudfront.net/collections_api_cios_public.html) |
|Contract API| [contract.openapi.yml](./openapis/contract.openapi.yml) | 0.0.1| [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/contract/api/) |
|Device API| [device.openapi.yml](./openapis/device.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/device/) |
|Device Asset API| [deviceasset.openapi.yml](./openapis/deviceasset.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/device_asset_lifecycle/) |
|File Storage API| [filestorage.openapi.yml](./openapis/filestorage.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/file_storage/) |
|Geography API| [geography.openapi.yml](./openapis/geography.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/geo_point/) |
|Group API| [group.openapi.yml](./openapis/group.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/group/) |
|License API| [license.openapi.yml](./openapis/license.openapi.yml) | 0.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/license/) |
|Pub/Sub API| [publishsublish.openapi.yml](openapis/pubsub.openapi.yml) | 1.0.0 |[chennel](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/pubsub/) / [messaging](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/pubsub_messaging/) / [datastore](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/pubsub_datastore/) |
|Resource Owner API| [resourceowner.openapi.yml](./openapis/resourceowner.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/resource_owner/) |
|Video Streaming API| [videostreaming.openapi.yml](./openapis/videostreaming.openapi.yml) | 1.0.0 | [portal](https://cios.gitlab.tokyo.optim.co.jp/developer-portal/cios-public/md/api_v2/video_streaming/) |
