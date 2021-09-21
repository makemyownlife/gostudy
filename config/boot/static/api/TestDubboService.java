package {{.BasePackage}}.api;

import {{.BasePackage}}.api.dto.HelloResult;

public interface TestDubboService {

    HelloResult hello(String mylife);

}
