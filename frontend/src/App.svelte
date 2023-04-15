<script lang="ts">
  import {ReceiveMessage, SendMessage} from '../wailsjs/go/main/App.js'
  import { onMount } from 'svelte';

  let Messages: string[] = ["Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi","Hi"]

  // async function receive() {
  //   while (true) {
  //     let result = await ReceiveMessage();

  //     if (result != '') {
  //       Messages = [...Messages, result];
  //       console.log(`Received message: ${result}`);
  //       console.log(`Messages: ${JSON.stringify(Messages)}`);
  //     }
  //   }
  // }

  async function receive() {
    async function recursiveReceive() {
      let result = await ReceiveMessage();
      if (result != '') {
        Messages = [...Messages, result];
        console.log(`Received message: ${result}`);
        console.log(`Messages: ${JSON.stringify(Messages)}`);
      }
      await recursiveReceive(); // 调用自身以进行下一次接收
    }
    await recursiveReceive(); // 首次调用
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

  // 滚动页面到最底部
  // let messageList = document.getElementById('messageList') as HTMLDivElement
  // messageList.scrollTop = messageList.scrollHeight;

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