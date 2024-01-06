package main

func main() {

}

// ls -la

// IMPORTANT
// permissions
// owner, group, world
// r, w, x
// 4, 2, 1 		(7)

/*
d = directory
rwxrwxrwx = owner, group, world
example:
drwxr-xr-x+		spockmcleod		staff
rwx = READ, WRITE, EXECUTE

d(rwx)(r-x)(r-x)
owner group world
the owner can READ, WRITE, EXECUTE
the group can READ, EXECUTE
the world can READ, EXECUTE

example:
chmod 777 temp.txt
7 = r, w, x (can do all 3)
so 777 means
the owner can READ, WRITE, EXECUTE
the group can READ, WRITE, EXECUTE
the world can READ, WRITE, EXECUTE

my example:
chmod 754 temp.txt
the owner can READ, WRITE, EXECUTE
the group can READ, EXECUTE
the world can READ
*/

// ^C = ctrl + c (when using nano *editing file*)
