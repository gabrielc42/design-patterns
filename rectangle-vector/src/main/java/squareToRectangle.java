class Square
{
    public int side;

    public Square(int side)
    {
        this.side = side;
    }
}

interface Rectangle
{
    int getWidth();
    int getHeight();

    default int getArea()
    {
        return getWidth() * getHeight();
    }
}

class SquareToRectangleAdapter implements Rectangle
{
    Square square;
    public SquareToRectangleAdapter(Square square)
    {
        this.square = square;
    }

    @Override
    public int getWidth() {
        return this.square.side;
    }

    @Override
    public int getHeight() {
        return this.square.side;
    }
}

 public class squareToRectangle {
     public static void main(String[] args) {
         SquareToRectangleAdapter squareToRectangleAdapter = new SquareToRectangleAdapter(new Square(5));
         System.out.println(squareToRectangleAdapter.getArea());
     }
 }