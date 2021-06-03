# Device

## Device Management API

### interface
```
GetDevices(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) (cios.MultipleDevice, *_nethttp.Response, error)
GetDevicesAll(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
GetDevicesUnlimited(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
GetDevice(ciosctx.RequestCtx, string, *string, *bool) (cios.Device, *_nethttp.Response, error)
GetDeviceInventory(ciosctx.RequestCtx, string) (map[string]interface{}, *_nethttp.Response, error)
DeleteDevice(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
CreateDevice(ciosctx.RequestCtx, cios.DeviceInfo) (cios.Device, *_nethttp.Response, error)
UpdateDevice(ciosctx.RequestCtx, string, cios.DeviceUpdateRequest) (cios.Device, *_nethttp.Response, error)
```

### Usage

#### Get a Device
```go
device, httpResponse, err := client.DeviceManagement.GetDevice(ctx, "device_id")
```

#### Get a Device Inventory
```go
inventory, httpResponse, err := client.DeviceManagement.GetDeviceInventory(ctx, "device_id")
```

#### Get Devices max limit 1000
```go
options := srvdevice.MakeGetDevicesOpts
devices, httpResponse, err := client.DeviceManagement.GetDevices(ctx, options())
```

#### Get Devices max no limit

```go
options := srvdevice.MakeGetDevicesOpts
devices, httpResponse, err := client.DeviceManagement.GetDevicesAll(ctx, options())
```

#### Get Devices max unlimited

```go
options := srvdevice.MakeGetDevicesOpts
devices, httpResponse, err := client.DeviceManagement.GetDevicesUnlimited(ctx, options())
```

### Create a Device 

```go
device, httpResponse, err := client.DeviceManagement.CreateDevice(ctx, cios.DeviceInfo{})
```

### Update a Device

```go
device, httpResponse, err := client.DeviceManagement.UpdateDevice(ctx, "device_id", cios.DeviceUpdateRequest{})
```

### Delete a Device

```go
httpResponse, err := client.DeviceManagement.DeleteDevice(ctx, "device_id")
```

## Device Monitoring API

### interface
```
GetMonitoringLatestList(ciosctx.RequestCtx, []string) ([]cios.DeviceMonitoring, *_nethttp.Response, error)
GetMonitoring(ciosctx.RequestCtx, string) (cios.DeviceMonitoring, *_nethttp.Response, error)

```

### Usage

#### Get Devices Monitoring info

```go
monitoringList, httpResponse, err := client.DeviceManagement.GetMonitoringLatestList(ctx, []string{"device_id1", "device_id2"})
```

#### Get a Device Monitoring info

```go
monitoring, httpResponse, err := client.DeviceManagement.GetMonitoring(ctx, "device_id1")
```


## Device Policy API

### interface
```
GetPolicies(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) (cios.MultipleDevicePolicy, *_nethttp.Response, error)
GetPoliciesAll(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
GetPoliciesUnlimited(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
DeletePolicy(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
CreatePolicy(ciosctx.RequestCtx, string) (cios.DevicePolicy, *_nethttp.Response, error)
```

### Usage

#### Get Policies max limit 1000

```go
options := srvdevice.MakeGetPoliciesOpts()
policies, httpResponse, err := client.DeviceManagement.GetPolicies(ctx, options())
```

#### Get Policies max no limit

```go
options := srvdevice.MakeGetPoliciesOpts()
policies, httpResponse, err := client.DeviceManagement.GetPoliciesAll(ctx, options())
```

#### Get Policies unlimited

```go
options := srvdevice.MakeGetPoliciesOpts()
policies, httpResponse, err := client.DeviceManagement.GetPoliciesUnlimited(ctx, options())
```

#### Create a Policy

```go
policy, httpResponse, err := client.DeviceManagement.CreatePolicy(ctx, "resource_owner_id")
```

#### Delete a Policy

```go
httpResponse, err := client.DeviceManagement.DeletePolicy(ctx, "policy_id")
```

