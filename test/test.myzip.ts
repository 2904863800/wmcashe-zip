import { MyZip } from "../src";

(async () => {
    try {
        const files = [
            "assets/ChromeSetup.exe",
            "assets/DiscordSetup.exe",
            "assets/WeChatSetup.exe",
            "assets/conf",
        ];
        const myZip = new MyZip();
        await myZip.zipFiles(files, "bfchain.zip", false);
    } catch (error) {
        console.log(error);
    }
})();
