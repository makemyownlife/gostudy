package {{.BasePackage}}.provider;

import com.alibaba.fastjson.JSON;
import {{.BasePackage}}.api.TestDubboService;
import {{.BasePackage}}.api.dto.DubboRpcResult;
import {{.BasePackage}}.domain.po.User;
import {{.BasePackage}}.service.UserService;
import org.apache.dubbo.config.annotation.Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

//属于Dubbo的@Service注解，非Spring  作用：暴露服务
@Service
@Component
public class TestDubboServiceImpl implements TestDubboService {

    private final static Logger logger = LoggerFactory.getLogger(TestDubboServiceImpl.class);

    @Autowired
    TestService testService;

    @Override
    public DubboRpcResult hello(String mylife) {
        return DubboRpcResult.custom(DubboRpcResult.SUCCESS, null);
    }

}
