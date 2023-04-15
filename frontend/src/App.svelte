<script lang="ts">
  import {ReceiveMessage, SendMessage} from '../wailsjs/go/main/App.js'
  import { onMount } from 'svelte';

  let Messages: string[] = ["Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi"]

  async function receive() {
  while (true) {
    let result = await ReceiveMessage();

    if (result != '') {
      Messages = [...Messages, result];
      console.log(`Received message: ${result}`);
      console.log(`Messages: ${JSON.stringify(Messages)}`);
    }
    }
  }

  receive()
  
  function sendMessage() {
    let input = document.getElementById('message-input') as HTMLInputElement
    if (input.value != ""){
      SendMessage(input.value)
    }
  }

  // input标签自动聚焦
  let inputElement: { focus: () => void; };
  onMount(() => {
    inputElement.focus();
  });

</script>
<div id="message-container">
  <div class="message-list">
    {#each Messages as message, i}
      <p class="message">
        <span> { message } </span>
      </p>
    {/each}
  </div>
</div>
<form id="send-container">
  <input type="text" id="message-input"  placeholder="请输入..." autocomplete="off"  bind:this={inputElement} />
  <button type="submit" id="send-button" on:click={sendMessage}>Send</button>
</form>

<!-- 默认 -->
<!-- <script lang="ts">
  import logo from './assets/images/logo-universal.png'
  im
  static push(result: string) {
    throw new Error('Method not implemented.');
  }port {Greet} from '../wailsjs/go/main/App.js'

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