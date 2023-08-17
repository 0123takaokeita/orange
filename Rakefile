############
##  関数  ##
############
def env_sh(cmd)
  sh "bash -c 'source bin/env.sh && #{cmd}'"
end

def de(args)
  sh %(docker exec #{args})
end

def dc(args)
  sh %(docker compose #{args})
end

# コンテナのステータスを確認
# @param [String | Symbol]
# @return [String]
def get_container_status(*container_names)
  a = container_names.map { |n| `docker inspect --format='{{.State.Status}}' #{n} 2>/dev/null`.chomp }
  a = a.first if a.length == 1
  a
end

############
##  task  ##
############
namespace :cmp do
  desc 'コンテナ立ち上げ'
  task :up do
    dc 'up -d --build'
  end

  desc 'コンテナを落とす'
  task :down do
    dc 'down'
  end
end

namespace :db do
  desc 'dbコンテナに入ってコンソールを開く'
  task :cui do
    unless get_container_status(:mariadb) == 'running'
      Rake::Task['cmp:up'].invoke
      sleep 0.5
    end

    env_sh %(docker exec -it $MYSQL_DATABASE mysql -u$MYSQL_USER -p$MYSQL_PASSWORD)
  end

  desc 'dbコンテナに入る'
  task :bash do
    Rake::Task['cmp:up'].invoke unless get_container_status(:mariadb) == 'running'
    env_sh %(docker exec -it $MYSQL_DATABASE bash)
  end

  desc 'restert'
  task :restart do
    sh 'docker restart mysql8'
  end

  desc 'DB backup'
  task :backup, [:dbname] do |_t, arg|
    env_sh %(docker exec $MYSQL_DATABASE mysqldump -u$MYSQL_USER -p$MYSQL_PASSWORD #{arg[:dbname]} > ./backup/#{Time.now.strftime('%Y%m%d%H%M%S')}_#{arg[:dbname]}_backup.sql)
  end

  desc 'DB restore'
  task :restore, [:dbname] do |_t, _arg|
    puts 'todo'
    # env_sh %(docker exec $MYSQL_DATABASE mysqldump -u$MYSQL_USER -p$MYSQL_PASSWORD #{arg[:dbname]} > ./sample/#{Time.now.strftime("%Y_%H_%M_%S")}_backup.sql)
  end
end
