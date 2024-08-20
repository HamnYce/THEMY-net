import fs from "fs";
import path from "path";

export function getOrderedFolders() {
  const appDir = path.join(process.cwd(), "src/app");
  const folders = fs.readdirSync(appDir).filter((file) => {
    const pagePath = path.join(appDir, file, "page.tsx");
    return fs.statSync(path.join(appDir, file)).isDirectory() && fs.existsSync(pagePath);
  });

  const order = [
    { name: "server", index: 1 },
    { name: "scanner", index: 2 },
    { name: "history", index: 3 },
    { name: "testBackend", index: 3 },
  ];

  const exclude = ["login","verifyEnv"];
  let defaultIndex = order.length + 1;
  const orderedFolders = folders
    .filter((folder) => !exclude.includes(folder))
    .map((folder) => {
      const orderEntry = order.find((o) => o.name === folder);
      if (orderEntry) {
        return { name: folder, index: orderEntry.index };
      } else {
        return { name: folder, index: defaultIndex++ };
      }
    })
    .sort((a, b) => a.index - b.index);

  return orderedFolders;
}
