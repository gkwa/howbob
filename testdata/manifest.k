{packages = [
    {
        name = "postgresql"
        check_installed = "postgres --version"
    }
    {
        name = "sqlite3"
        check_installed = "sqlite3 --version"
    }
    {
        name = "mysql"
        check_installed = "mysqld --version"
    }
    {
        name = "lua"
        check_installed = "lua -v"
    }
    {
        name = "lua-language-server"
        check_installed = "find -L /home/linuxbrew/.linuxbrew/bin/lua-language-server"
    }
    {
        name = "cdk8s"
        check_installed = "cdk8s --version"
    }
]}
