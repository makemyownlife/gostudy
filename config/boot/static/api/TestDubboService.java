package {{.BasePackage}}.api;

import {{.BasePackage}}.api.dto.DubboRpcResult;

public interface TestDubboService {

    DubboRpcResult hello(String mylife);

}
