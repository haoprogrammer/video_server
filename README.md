# 流媒体网站


## 1.API设计：用户
<table>
  <tr>
    <th>功能</th>
    <th>RESTful 风格URL</th>
    <th>METHOD</th>
    <th>SC(状态码)</th>
  </tr>
  <tr>
    <td>创建（注册用户）</td>
    <td>URL:/user </td>
    <td> POST </td>
    <td> 201,400,500 </td>
  </tr>
  <tr>
    <td>用户登录</td>
    <td> URL:/user/:username </td>
    <td> POST </td>
    <td> 200,400,500 </td>
  </tr>
  <tr>
    <td>获取用户基本信息</td>
    <td> URL:/user/:username </td>
    <td> GET </td>
    <td> 200,400,401,403,500 </td>
  </tr>
  <tr>
    <td>用户注销</td>
    <td> URL:/user/:username </td>
    <td> DELETE </td>
    <td> 200,400,401,403,500 </td>
  </tr>

</table>