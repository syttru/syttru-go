package main

import(
  "fmt";
  "io";
  "os";
)

func main(){
  // パラメータをチェック
  if len(os.Args) <= 1 {
    fmt.Println("ファイルを指定してください");
    return;
  }

  // ファイルを読み込み
  data, err := io.ReadFile(os.Args[1]);
  if err != nil {
    fmt.Println(err);
    return;
  }

  // 出力
  fmt.Println(data);
}

