function encode(list: string[]): string {
  let output = "";

  for (const str of list) {
    output += str.length + "#" + str;
  }

  return output;
}

function decode(str: string): string[] {
  const output: string[] = [];
  let i = 0;

  while (i < str.length) {
    let j = i;
    while (str[j] !== "#") {
      j++;
    }

    let length = Number(str.slice(i, j));
    let chunk = str.slice(j + 1, j + 1 + length);
    output.push(chunk);

    i = j + 1 + length;
  }

  return output;
}

console.log('encode', encode(["lint","code","love","you"]));
const inp = encode(["lint","code","love","you"]);
console.log('decode', decode(inp));
