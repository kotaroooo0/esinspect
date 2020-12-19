FROM openjdk:8-alpine as dict_builder

WORKDIR /home

# Sudachiプラグインのjarファイルを持ってくる (バイナリ辞書の作成のため)
RUN wget https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v7.8.1-2.0.3/analysis-sudachi-7.8.1-2.0.3.zip && \
    unzip analysis-sudachi-7.8.1-2.0.3.zip && \
    # 用意されているシステム辞書を持ってくる
    wget https://object-storage.tyo2.conoha.io/v1/nc_2520839e1f9641b08211a5c85243124a/sudachi/sudachi-dictionary-20200722-core.zip && \
    unzip sudachi-dictionary-20200722-core.zip

## ユーザー辞書ソースを持ってくる
COPY sudachi/custom_dict.csv /home
# バイナリ辞書の作成
RUN java -Dfile.encoding=UTF-8 -cp /home/sudachi-0.4.3.jar com.worksap.nlp.sudachi.dictionary.UserDictionaryBuilder -o /home/custom.dic -s /home/sudachi-dictionary-20200722/system_core.dic /home/custom_dict.csv


FROM elasticsearch:7.8.1

# Sudachiプラグインの設定ファイル
COPY sudachi/sudachi.json /usr/share/elasticsearch/config/sudachi/
# 類義語辞書を追加
COPY sudachi/synonyms.txt /usr/share/elasticsearch/config/sudachi/
# 前ステージでダウンロードしたSudachiのシステム辞書
COPY --from=dict_builder /home/sudachi-dictionary-20200722/system_core.dic /usr/share/elasticsearch/config/sudachi/
# 前ステージで作ったユーザー辞書
COPY --from=dict_builder /home/custom.dic /usr/share/elasticsearch/config/sudachi/
# 前ステージでダウンロードしたプラグイン
COPY --from=dict_builder /home/analysis-sudachi-7.8.1-2.0.3.zip /usr/share/elasticsearch/

# Sudachiプラグインインストール
RUN elasticsearch-plugin install file:///usr/share/elasticsearch/analysis-sudachi-7.8.1-2.0.3.zip && \
    elasticsearch-plugin install analysis-icu && \
    rm /usr/share/elasticsearch/analysis-sudachi-7.8.1-2.0.3.zip
