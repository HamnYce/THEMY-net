import fs from "fs";
import path from "path";

export function getOrderedFolders() {
  const appDir = path.join(process.cwd(), "src/app");
  const folders = fs.readdirSync(appDir).filter((file) => {
    const pagePath = path.join(appDir, file, "page.tsx");
    return fs.statSync(path.join(appDir, file)).isDirectory() && fs.existsSync(pagePath);
  });

  const order = [
    { name: "dashboard", index: 1 },
    { name: "networkAnalysis", index: 2 },
    { name: "networkReports", index: 3 },
    { name: "security", index: 4 },
    { name: "testBackend", index: 5 },
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
