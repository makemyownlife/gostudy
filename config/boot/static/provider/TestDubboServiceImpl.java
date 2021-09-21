package {{.BasePackage}}.provider;

import com.alibaba.fastjson.JSON;
import {{.BasePackage}}.api.TestDubboService;
import {{.BasePackage}}.api.dto.HelloResult;
import {{.BasePackage}}.domain.po.User;
import {{.BasePackage}}.service.UserService;
import org.apache.dubbo.config.annotation.Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Service //属于Dubbo的@Service注解，非Spring  作用：暴露服务
@Component
public class TestDubboServiceImpl implements TestDubboService {

    private final static Logger logger = LoggerFactory.getLogger(TestDubboServiceImpl.class);

    @Autowired
    UserService userService;

    @Override
    public HelloResult hello(String mylife) {
        HelloResult result = new HelloResult();
        User user = userService.getUserById(1L);
        logger.info("user:" + JSON.toJSONString(user));
        result.setResult("张勇" + mylife);
        return result;
    }

}