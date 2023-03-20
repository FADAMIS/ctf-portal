<script>
    //send login request with fetch
    let username
    let password
    function login() {
        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        }).then(res => {
            if (res.status == 200) {
                window.location.href = '/dashboard'
            } else {
                alert('Wrong username or password')
            }
        })
    }
</script>
<main>
    <div id="body" class="grid w-screen h-screen place-content-center"> 
        <div class="w-64">
        <h1 class="text-gray-300 text-4xl text-center font-bold mb-3 font-mono">CTF Login</h1>
            <div class="grid place-items-center">
                <label class="text-gray-300 font-mono mb-1" for="username">team_name</label>
                <input class="rounded-lg h-9 w-full p-2 transition-all bg-black hover:bg-violet-900 text-gray-300" type="text" name="username" id="username" bind:value={username}/>
                <label class="text-gray-300 font-mono mt-2 mb-1" for="password">password</label>
                <input class="rounded-lg h-9 w-full p-2 transition-all bg-black hover:bg-violet-900 text-gray-300" type="password" name="password" id="password" bind:value={password}/>
                <button class="bg-gray-900 w-32 h-12 rounded-xl transition-all border-2 hover:border-4 border-violet-500 mt-2" type="submit" on:click={login}><p class="text-gray-300 font-mono">Login</p></button>
            </div>

        </div>
    </div>
</main>
<style>
    #body {
        background-color: #161616;
        opacity: 0.8;
        background-size: 10px 10px;
        background-image: repeating-linear-gradient(45deg, #2e2e2e 0, #2e2e2e 1px, #040404 0, #010101 50%);
    }
</style>