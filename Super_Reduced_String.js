//
function processData(str) {
  const ans = solve(Array.from(str));
  if (!ans.length) console.log("Empty String");
  else {
    console.log(ans.join(""));
  }
}

function solve(data) {
  const [ajusted, ret] = step(data);
  if (ajusted) {
    return solve(ret);
  }
  return ret;
}

function step(arr) {
  let i, len;
  let ajusted = false;
  const before = [];
  for(i = 1, len = arr.length; i < len;) {
    if (arr[i] !== arr[i - 1]) {
      before.push(arr[i - 1]);
      i++;
    } else {
      ajusted = true;
      break;
    }
  }
  if (ajusted){
    return [ajusted, before.concat(arr.slice(i + 1))];

  }
  return [ajusted, arr];
}

processData("aab");
processData("baab");