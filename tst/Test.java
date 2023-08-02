public class Test {
  private String foo;
  public Test() {
    this.foo = "Yes";
  }
    public static void main(String[] args) {
      Test test = new Test();
      test.foo();
    }

    public boolean foo() {
      return false;
    }
}

