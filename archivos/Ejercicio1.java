package Recursion;

public class Ejercicio1 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
			
		int matriz[][] = new int[3][3];
		
		matriz[0][0] = 2;
		matriz[0][1] = 4;
		matriz[0][2] = 4;
		matriz[1][0] = 6;
		matriz[1][1] = 6;
		matriz[1][2] = 9;
		matriz[2][0] = 8;
		matriz[2][1] = 10;
		matriz[2][2] = 12;
		
		System.out.print("La suma de la matriz es: " + sumaM(matriz, 0, 0));
		
		
	}

	
   public static  int sumaM(int[][]mat, int i, int j){
	   
	   System.out.println(mat[i][j] + " ");
	   
	   if(i!= mat.length-1 || j!= mat.length-1) 
	   {
		   if(j==mat[i].length-1) {
			   
			   return mat[i][j] + sumaM(mat, i+1, 0);
			   
		   }else {
			   
			   return mat[i][j] + sumaM(mat, i, j+1);
		   }
		   
		   
		   
	   }
	   

	   return mat[i][j];
	   
       
	   
	   
    }
}
