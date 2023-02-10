package {{.BasePackage}}.service;

import {{.BasePackage}}.domain.mapper.TestMapper;
import {{.BasePackage}}.domain.po.TestPo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TestService {

    @Autowired
    private TestMapper TestMapper;

}
