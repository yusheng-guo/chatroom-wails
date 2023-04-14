<script lang="ts">
  let messages = [
    "hellowhat?what?what?what?what?what?what?what?what?what?what?what?what?what?what?what?what?",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
    "hello",
    "what?",
    "hi!",
  ];

  let messagesObject = [];
  const setListItemWidth = () => {
    const windowWidth = window.innerWidth;
    const maxWidth = windowWidth * 0.8; // 窗口尺寸的 80%
    messagesObject = messages.map(item => {
      const messageWidth = item.length * 20; // 字符长度乘以字体大小
      return {
        text: item,
        width: messageWidth < maxWidth ? messageWidth : maxWidth
      }
    });
  }
  window.addEventListener('resize', setListItemWidth); // 监听窗口大小变化

  setListItemWidth(); // 初始化每个列表项的宽度
</script>

<div id="message-container">
  <ul class="message-list">
    {#each messagesObject as message}
      <li class="message" style="width: {message.width}px;">{message.text}</li>
    {/each}
  </ul>
</div>
<form id="send-container">
  <input type="text" id="message-input"  placeholder="请输入...">
  <button type="submit" id="send-button">Send</button>
</form>

<!-- <div class="chat">
  <div class="chat-window-container">
    <div class="chat-window">
      {#each messages as message, i}
        <div class="message {message.sent ? 'sent' : 'received'}">
         <div class="message-text">{message.text}</div>
        </div>
      {/each}
    </div>
  </div>
  <div class="input-container">
    <div class="input">
      <input id="message-input" type="text" placeholder="Type a message..." />
      <button on:click={sendMessage}>Send</button>
    </div>
  </div>
</div> -->


<!-- 默认 -->
<!-- <script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {Greet} from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Greet</button>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style> -->
