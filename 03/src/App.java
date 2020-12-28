import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.stream.Collectors;

/**
 * This program is to solve AOc day 2
 * The Q is:
 * Starting at the top-left corner of your map 
 * and following a slope of right 3 and down 1, 
 * how many trees would you encounter?
 * @author Yee Heng
 * 
 */

public class App {
    static long countTree(String[] input, int dy, int dx) {

        final int height = input.length;
        final int width = input[0].length();

        int x = 0;
        int y = 0;
        long result = 0;

        while (y < height) {
            int y1 = y + dy;
            int x1 = (x + dx) % width;

            if (y1 < height) {
                if (input[y1].charAt(x1) == '#') {
                    result += 1;
                }
            }
            // update the coordinates
            x = x1;
            y = y1;
        }
        // System.out.println("*********\nPart 1*********\n\nCount:\t" + result);
        return result;
    }

    static long part1(String[] input){
        return countTree(input, 1, 3);
    }

    static long part2(String[] input){
        
        int slope[][]={
            // {x, y}
            {1, 1},
            {1, 3},
            {1, 5},
            {1, 7},
            {2, 1},
        };
        long result = 1;
        for (int i = 0; i < slope.length; i++) {
            result *= countTree(input, slope[i][0], slope[i][1]);
            System.out.println(countTree(input, slope[i][0], slope[i][1]));
        }
        return result;
    }

    public static void main(String[] args) throws IOException {
        String[] input = null; // initialize the input variable

        // read file and convert to an array of string
        try (var file = new BufferedReader(new FileReader("input.txt"))){
            input = file.lines().collect(Collectors.toList()).toArray(new String[0]);
        }

        System.out.printf("*********\nPart 1\n*********\nCount:\t%d\n\n",part1(input));
        System.out.printf("*********\nPart 2\n*********\nCount:\t%d\n\n",part2(input));

        // System.out.println("*********\nPart 2 Debug\n*********\nCount:\n");
        // part2_debug(input);
        // System.out.println(countTree(input, 1, 1));
        // System.out.println(countTree(input, 1, 3));
        // System.out.println(countTree(input, 1, 5));
        // System.out.println(countTree(input, 1, 7));
        // System.out.println(countTree(input, 2, 1));
    }
}   