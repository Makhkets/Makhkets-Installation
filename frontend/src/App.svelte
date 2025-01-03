
<script>
  import logo from './assets/images/logo-universal.png'
  import { writable } from 'svelte/store';
  import { onMount } from 'svelte';


  import { InstallApplications } from '../wailsjs/go/main/App.js'
  import { GreetinForUser, Minimize, Close } from '../wailsjs/go/main/App.js'

  import './assets/css/style.css';


  let requestResult = "";
  let inputText = "";



  let currentComponent = writable('App');
  let selectedTheme = writable("Darcula");

  // Реакция на выбор темы
  $: {
    if ($selectedTheme === "Darcula") {
      document.documentElement.style.setProperty("--bg-color", "#282a36");
      document.documentElement.style.setProperty("--bg-main-color", "#44475a");
      document.documentElement.style.setProperty("--button-active", "#6271a4");
      document.documentElement.style.setProperty("--text-color", "#d4d8e4");
      document.documentElement.style.setProperty("--username-color", "#fc7ac7");
      document.documentElement.style.setProperty("--menu-button-color", "#44475a");
      document.documentElement.style.setProperty("--menu-button-active", "#6271a4");
    } else {
      document.documentElement.style.setProperty("--bg-color", "#111827");
      document.documentElement.style.setProperty("--button-active", "#3B82F6");
      document.documentElement.style.setProperty("--text-color", "#b1bac0");
      document.documentElement.style.setProperty("--bg-main-color", "#06172b");
      document.documentElement.style.setProperty("--username-color", "#47bbf7");
      document.documentElement.style.setProperty("--menu-button-color", "#2D3748");
      document.documentElement.style.setProperty("--menu-button-active", "#3B82F6");
    }
  }

  function showApp1() {
    currentComponent.set('App')
  }
  function showApp2() {
    currentComponent.set('Contact');
  }



  function greetingForUser() {
    GreetinForUser(inputText).then(result => requestResult = result)
  }


  let progress = 0; // Текущий прогресс
  let selectedApplications = [];
  let isInstalling = writable(false); // Начало установки


  function handleCheckboxChange(event) {
    const { value, checked } = event.target;
    if (checked) {
      selectedApplications.push(value);
    } else {
      selectedApplications = selectedApplications.filter(app => app !== value);
    }
  }

  async function startInstallation() {
    if ($isInstalling) return; // Проверка на текущую активность
    if (selectedApplications.length === 0) {
      alert('Пожалуйста, выберите хотя бы одно приложение для установки.');
      return;
    }

    isInstalling.set(true);


    progress = 0;
    try {
      await InstallApplications(selectedApplications);
    } catch (error) {
      console.error('Ошибка при установке приложений:', error);
    }
  }

  function updateProgress(value) {
    progress = value;
    if (progress >= 100) {
      setTimeout(() => {
        progress = 0;
        isInstalling.set(false);
      }, 1500);
    }
  }

  // Регистрируем функцию в глобальной области видимости для вызова из Go
  onMount(() => {
    window.updateProgress = updateProgress;
  });



// function startInstallation() {
//     progress = 0; // Сброс прогресса перед началом
//     intervalId = setInterval(() => {
//       progress += 10;
//       if (progress >= 100) {
//         clearInterval(intervalId);
//         // Сбрасываем прогресс через 2 секунды
//         setTimeout(() => {
//           progress = 0;
//         }, 1500);
//       }
//     }, 500);
//   }
  /////

</script>


<main>
  <div class="window-controls">
    <button id="minimize-btn" on:click={Minimize}>_</button>
    <button id="close-btn" on:click={Close}>X</button>
  </div>

  <div class="sidebar">
    <div>
      <center><h1 class="username">MAKHKETS</h1></center>

      <div class="menu">
        <button
                on:click={showApp1}
                class:active={$currentComponent === 'App'}
        >
          Main
        </button>
        <button
                on:click={showApp2}
                class:active={$currentComponent === 'Contact'}
        >
          Contacts
        </button>
      </div>

      <!-- Theme Selector Dropdown -->
      <div class="theme-selector">
        <label for="theme-dropdown">Theme:</label>
        <select id="theme-dropdown" bind:value={$selectedTheme}>
          <option value="Makhkets">Makhkets</option>
          <option value="Darcula">Darcula</option>
        </select>
      </div>
    </div>

    <div class="footer">
      <a href="https://t.me/makhekts" target="_blank">Telegram</a>
    </div>
  </div>

  {#if $currentComponent === 'App'}
    <div class="content_app">
      <div class="content_app_main">
        <h2>Installation</h2>
        <p>-- Select the applications you want to install --</p>
      </div>

      <div class="container">
        <div class="checkbox-group">
          <label>
            <input on:change={handleCheckboxChange} type="checkbox" name="applications" value="goland">
            Goland Crack
          </label>
          <label>
            <input on:change={handleCheckboxChange} type="checkbox" name="applications" value="nekoray">
            Nekoray VPN
          </label>
        </div>
        <button
                class="install-button {$isInstalling ? 'disabled' : ''}"
                on:click={startInstallation}
                disabled={$isInstalling}>
          {$isInstalling ? `Installing ${progress}%` : 'Install'}

        </button>
      </div>



    </div>

  {:else if $currentComponent === 'Contact'}
    <div class="content_app">
      <h2>Contacts</h2>
      <p>waiting...</p>
    </div>
  {/if}

  <!-- Прогресс-бар -->
  <div class="progress-bar-container">
    <div class="progress-bar" style="width: {progress}%;"></div>
  </div>

</main>

<style>


  .sidebar {
    width: 250px;
    height: 100vh; /* Полная высота экрана */
    /*background-color: #282a36;*/

    background-color: var(--bg-color);

    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 20px;
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.5);
    position: fixed; /* Фиксируем сайдбар */
    top: 0;
    left: 0;
  }

  .username {
    font-size: 24px;
    margin-bottom: 20px;
    font-family: "Source Code Pro", monospace;
    /*color: #ff79c6;*/
    color: var(--username-color);
    letter-spacing: 1px;
  }

  .menu {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .menu button {
    padding: 10px;
    font-size: 16px;
    background-color: var(--menu-button-color);
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s;
    color: #A0AEC0;
  }

  .menu button:hover {
    /*background-color: #6272a4;*/
    background-color: var(--menu-button-active);
    color: #F3F4F6;
  }

  .menu button.active {
    color: #F3F4F6;
    /*background-color: #6272a4;*/
    background-color: var(--menu-button-active);
  }

  .theme-selector {
    margin-top: 20px;
    padding-left: 5px;
  }

  .footer {
    text-align: center;
  }

  .theme-selector label {
    font-size: 14px;
    margin-right: 5px;
  }

  select {
    background-color: var(--bg-color);
    color: var(--text-color);
    border: 1px solid var(--text-color);
    padding: 5px;
    border-radius: 5px;
  }

  .footer a {
    color: var(--green-color);
    text-decoration: none;
    font-size: 14px;
  }

  .footer a:hover {
    text-decoration: underline;
  }

  .content_app {
    margin-left: 255px; /* Сдвигаем основной контент, чтобы он не перекрывался с сайдбаром */
    padding: 20px;
    width: calc(100% - 250px); /* Чтобы контент занимал оставшуюся ширину */
    overflow-y: auto;

  }
</style>