package {{.BasePackage}}.domain.mapper;

import {{.BasePackage}}.domain.po.TestPo;
import org.springframework.stereotype.Repository;

@Repository
public interface TestMapper {

    TestPo getTestById(Long id);

}
