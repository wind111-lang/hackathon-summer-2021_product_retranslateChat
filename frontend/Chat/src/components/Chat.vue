<template>
    <div id = "app">
        <div class = "teamtag">#駆け出しエンジニアと繋がりたい</div>
    <div class = "backplace">
        
        <div class="title">
            再翻訳チャットルームだよ^_−☆ﾊﾞﾁｺｰﾝ<br>
        </div>
            <div id ="subtitle">
                日本語→英語→日本語です<br>(ユーザー名):「メッセージ」(時間)
            </div>
        <div class = "chatplace">
            <div class = "chatbox">
                ようこそ再翻訳チャットルームへ！<br>
                <div v-for="(message,key) in message" :key="key">
                    <div class="chat">{{message.message}}</div>
                </div>
            </div>
        </div>
        <section>
        <div class = "sentform">
            <input v-model="name" class="name"  placeholder="ユーザー名を入力">
            <textarea v-model="text" class="text"   placeholder="メッセージを入力"></textarea>
            <butto class="btn btn--orange btn--cubic btn--shadow" v-on:click="sendMessage">送信する</butto>
        </div>
        </section>
    </div>
    </div>
</template>

<script>


export default{
    data:function(){
        return{

        ws:null,
        message: [],
        text:"",
        time:"00:00:00",
        name:[],
        // serverUrl:'ws://localhost:8081/ws'
        
    }
    },
    watch: {
        name(name) {
            this.name = this.charaLimit(name);  //名前の監視
    }
    },
    mounted:function() {
        this.connectToWebsocket();
    },
    
    methods:{
        
        connectToWebsocket() {
        var url = 'ws://' + window.location.host + '/ws'
        this.ws = new WebSocket(url);
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        this.ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
        },
        onWebsocketOpen() {
        console.log("connected to WS!");
        },
        sendMessage(){
            if(this.text !== ""){

                let SendMes = this.text
                let SendNam = this.name
                this.ws.send(JSON.stringify({message: SendMes,name: SendNam}));
                this.text="";
            }
        },
        async handleNewMessage(event){
            let response = event.data;
            let data = JSON.parse(response)
            // data = data.split(/\r?\n/);
            let date = new Date();  //時間の取得
            // this.time = date.getHours() + ":"+ date.getMinutes() + ":" +date.getSeconds();
            var hour =this.set2fig(date.getHours())
            var min =this.set2fig(date.getMinutes())
            var sec =this.set2fig(date.getSeconds())
            this.time = hour + ":" + min + ":" + sec
            
            let msg = data.text;
            let nam = data.name;
            let fullmsg = nam + ":「" + msg + "」(" + this.time + ")"
            this.message.push({message:fullmsg});
            this.scrollBotom();
            
            // var chatplace = this.$el.querySelector("#chatplace");
            // container.scrollTop = container.scrollHeight;
        },
        set2fig(num) {
        // 桁数が1桁だったら先頭に0を加えて2桁に調整する
        var ret;
        if( num < 10 ) { ret = "0" + num; }
        else { ret = num; }
        return ret;
},

        
        charaLimit(name) {
        return name.length > 256 ? name.slice(0, -1) : name;    //名前の上限は256
        },
        scrollBotom: function() {       //チャットの下部に移動
        var chatplace = this.$el.querySelector(".chatplace");
        chatplace.scrollTop = chatplace.scrollHeight;
    },
    
}
    
}


</script>

<style>
#app{
    margin: 0px;

}

.teamtag{
    float: left;
    font-size: 10px;

}

.title {
    font-size: 40px;
    padding: 0.5em;/*文字周りの余白*/
    color: #010101;/*文字色*/
    background: #eaf3ff;/*背景色*/
    border-bottom: solid 3px #516ab6;/*下線*/
}

.subtitle{
    font-size: 10px;
    padding: 0.5em;/*文字周りの余白*/
    color: #010101;/*文字色*/
    background: #eaf3ff;/*背景色*/
    border-bottom: solid 3px #516ab6;/*下線*/
}

.backplace{
    height: 400px;
    width: 100%;
}

.chatbox{
    font-size: 20px;
    clear: both;
}

.chatplace{
    height: 100%;
    width: 100%;
    border: 2px solid;
    border-color: blueviolet;
    overflow: scroll;
    margin: 2px;
    background: linear-gradient(-135deg, #9941D8 , #E4A972) fixed;
}

.chat{
    float: left;
    clear: both;
}

.sentform{
    border-top:2px aqua solid;
    
    height: 72px;
    width: 100%;
    margin: 8px;
    padding-top: 8px;
}
.name{
    float: left;
    width: 10%;
    margin-right: 2px;
}
.text{
    width: 70%;
    height: 36px;
    margin-left: 2px;
    float: left;
}
.btn{
    clear: both;
    size:100;
}
.btn--orange {
    color: #fff;
    background-color: #eb6100;
    border-bottom: 5px solid #b84c00;
}
.btn--orange:hover {
    cursor: pointer;
    margin-top: 3px;
    color: #fff;
    background: #f56500;
    border-bottom: 2px solid #b84c00;
}
.btn--shadow {
    -webkit-box-shadow: 0 3px 5px rgba(0, 0, 0, .3);
    box-shadow: 0 3px 5px rgba(0, 0, 0, .3);
}
</style>
