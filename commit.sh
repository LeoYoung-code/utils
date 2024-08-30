 # 检查是否传入了提交信息
  if [ -z "$1" ]; then
    echo "请提供提交信息"
    exit 1
  fi

  # 添加更改并提交
  git add . && git commit -m "$@"

  # 获取最新的提交哈希值
  commit_hash=$(git rev-parse HEAD)

  # 定义颜色和样式
  RED_BOLD='\033[1;31m'
  RESET='\033[0m'

  # 显示被删除的代码，并替换前缀，应用颜色样式
  git show $(git rev-parse HEAD) -- ':!*.pb.go' |  awk -v RED_BOLD="$RED_BOLD" -v RESET="$RESET" '/^-/ {getline next_line; if (next_line !~ /^\+/) print RED_BOLD "您删除了: " RESET substr($0, 2)}'


  # git diff
#  git diff | awk -v RED_BOLD="$RED_BOLD" -v RESET="$RESET" '/^-/ {getline next_line; if (next_line !~ /^\+/) print RED_BOLD "您删除了: " RESET substr($0, 2)}'