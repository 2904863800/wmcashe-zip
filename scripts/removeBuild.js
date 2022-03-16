// @ts-check

const { deleteFolder } = require("@wmcashe/devkits");

function removeBuild() {
    const sourcePath = process.cwd() + "/build";
    deleteFolder(sourcePath);
    console.log(`移除 ${sourcePath} 完成`);
}

removeBuild();
