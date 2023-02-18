class AdapterSquare
{
    public int side;

    public AdapterSquare(int side)
    {
        this.side = side;
    }
}

interface AdapterRectangle
{
    int getWidth();
    int getHeight();

    default int getArea()
    {
        return getWidth() * getHeight();
    }
}

class AdapterSquareToRectangleAdapter implements Rectangle
{
    Square square;
    public AdapterSquareToRectangleAdapter(Square square)
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

// public class SquareToRectangleAdapterDemo {
//     public static void main(String[] args) {
//         SquareToRectangleAdapter squareToRectangleAdapter = new squareToRectangleAdapter(new Square(5));
//         System.out.println(SquareToRectangleAdapter.getArea());
//     }
// }