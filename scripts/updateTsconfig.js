// @ts-check

const fs = require("fs");
const { listAllFile } = require("@wmcashe/devkits");
const runPath = process.cwd();

function updateTsConfig(sourcePath) {
    const filePathList = listAllFile(sourcePath, true, ["ts"]);
    const tsconfigPath = `${sourcePath}/tsconfig.json`;
    if (!fs.existsSync(tsconfigPath)) {
        throw new Error(`tsconfig.json 不存在 ${tsconfigPath}`);
    }
    const config = JSON.parse(fs.readFileSync(tsconfigPath).toString());
    config.files = filePathList;
    fs.writeFileSync(tsconfigPath, JSON.stringify(config, null, 4));
    console.log(`更新 ${sourcePath} 的 tsconfig.json 完成`);
}

const dirs = ["src", "test"];
dirs.map((dir) => updateTsConfig(`${runPath}/${dir}`));
