package rpc_objects

type Args struct{
	N,M int
}
func (a *Args)Multiply(args *Args,reply *int)error{
	*reply=args.M*args.N
	return nil
}
