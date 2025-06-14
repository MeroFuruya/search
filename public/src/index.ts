function helloWorld(n: number = 1) {
  console.log("Hello, World!");
  if (n > 1) {
    return helloWorld(n - 1);
  }
}

helloWorld(10);

console.log("This is the main entry point of the application.");