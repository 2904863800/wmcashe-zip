import * as os from "node:os";
import * as path from "node:path";
import * as childProcess from "node:child_process";

export class MyZip {
    zipFiles(files: string[], output: string, isCompress = false) {
        return new Promise((resolve, reject) => {
            const platform = os.platform();
            const suffix = platform === "win32" ? ".exe" : "";
            const ziperPath = path.join(__dirname, `bin/myzip${suffix}`);

            let cmd = `${ziperPath} -files ${files.join(",")} -output ${output}`;

            if (isCompress) {
                cmd += ` -isCompress`;
            }

            childProcess.exec(cmd, (error, stdout, stderr) => {
                if (error) {
                    return reject(error);
                }
                console.log(stdout);
                console.log(`compress success`);
                resolve(stdout);
            });
        });
    }
}
