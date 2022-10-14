## APIHut

[![Build Status](https://drone.northes.co/api/badges/apihut/server-2.0/status.svg?ref=refs/heads/main)](https://drone.northes.co/apihut/server-2.0)

### 部署流程
1. 创建基础镜像 `deploy/images`
2. 设置 pvc，同步到 `.drone.yml` 的 `volumes` 项中
3. harbor 创建机器人账户，拥有镜像推送权限（如果使用了预构建的镜像，创建的账户需要拥有FROM镜像的拉取权限），设置到 drone 的环境变量
   `docker_username` `docker_password`
4. drone 项目开启 `Trusted` 和 `Auto cancel pull requests` 和 `Auto cancel running`