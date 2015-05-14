using Microsoft.Xna.Framework;
using mopacman.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Controllers
{
    class GhostAIController : GameComponent
    {
        public Ghost Ghost { get; private set; }
                
        public IControllable Target { get; private set; }

        public GhostAIController(MyGame g, Ghost player, IControllable target)
            : base(g)
        {
            this.Ghost = player;
            this.Target = target;
        }

        public override void Update(GameTime gameTime)
        {
            //Modo perseguição sempre avança em busca do alvo
            if (this.Ghost.State == Ghost.States.Chase)
            {
                this.scatterNextDestination = null;

                ChaseSection(this.Target.CurrentLocation);
            }
            else if (this.Ghost.State == Components.Ghost.States.Scatter)
            {
                if (this.scatterNextDestination == null)
                {
                    this.scatterNextDestination = this.Ghost.Region.Item1;
                }
                else if (this.scatterNextDestination.ID == this.Ghost.CurrentLocation.ID)
                {
                    this.scatterNextDestination = (this.scatterNextDestination.ID == this.Ghost.Region.Item1.ID) ? this.Ghost.Region.Item2 : this.Ghost.Region.Item1;
                }

                this.ChaseSection(this.scatterNextDestination);
            }
        }

        protected void ChaseSection(MazeSection destination)
        {
            //Localizações dos elementos 
            var currLocation = this.Ghost.CurrentLocation;
            var currEndLocation = destination;
            var prevLocation = this.Ghost.PreviousLocation;

            //Essa é a seção inicial, 
            PathFindingNode currNode = new PathFindingNode();
            currNode.h = 0;
            currNode.g = 0;
            currNode.parent = null;
            currNode.location = currLocation;

            List<PathFindingNode> closedSet = new List<PathFindingNode>();
            List<PathFindingNode> openSet = new List<PathFindingNode>();

            closedSet.Add(currNode);

            do
            {
                //Percorrendo as opções de caminho que temos
                List<MazeSection> opcoes = new List<MazeSection> { currLocation.W, currLocation.N, currLocation.E, currLocation.S };

                foreach (var o in opcoes)
                {

                    if (o != null)
                    {

                        //A seção precisa estar ativa e não ser a seção anteriormente visitada
                        if (o.Allowed && (prevLocation == null || o.ID != prevLocation.ID))
                        {
                            //Seção não deve existir na lista de seções fechadas
                            if (!closedSet.Any(x => x.location.ID == o.ID))
                            {
                                //...nem na lista de seções abertas
                                if (!openSet.Any(x => x.location.ID == o.ID))
                                {

                                    PathFindingNode node = new PathFindingNode();
                                    node.location = o;
                                    node.parent = currNode;
                                    node.h = this.Heuristics(node.location, currEndLocation); //heuristica manhattan

                                    //Ok, esse nó participará do teste 
                                    openSet.Add(node);
                                }
                            }
                        }
                    }
                }

                if (openSet.Count == 0)
                {
                    break;
                }

                //Procuramos o melhor nó localizado (custo)
                var minNode = openSet[0];

                foreach (var m in openSet)
                {
                    if (m.h < minNode.h)
                    {
                        minNode = m;
                    }
                }

                currNode = minNode;
                currLocation = currNode.location;

                //Removendo da lista de nós abertos para pesquisa e incluíndo na lista de nós proibidos
                openSet.Remove(minNode);
                closedSet.Add(currNode);

            } while (currNode.location.ID != currEndLocation.ID);


            if (currNode.location.ID == currEndLocation.ID)
            {
                ///*

                LinkedList<PathFindingNode> path = new LinkedList<PathFindingNode>();
                PathFindingNode lastPathNode = currNode;
                MazeSection nextSectionToMove = null;
                MazeSection mySection = this.Ghost.CurrentLocation;

                while (lastPathNode != null)
                {
                    path.AddFirst(lastPathNode);
                    lastPathNode = lastPathNode.parent;
                }

                nextSectionToMove = path.Count > 1 ? path.First.Next.Value.location : path.First.Value.location;

                if (nextSectionToMove.ID == mySection.N.ID)
                {
                    this.Ghost.GoTo(EnumDirections.North);
                }
                else if (nextSectionToMove == mySection.S)
                {
                    this.Ghost.GoTo(EnumDirections.South);
                }
                else if (nextSectionToMove == mySection.E)
                {
                    this.Ghost.GoTo(EnumDirections.East);
                }
                else if (nextSectionToMove == mySection.W)
                {
                    this.Ghost.GoTo(EnumDirections.West);
                }
                //*/
            }
        }

        private Int32 Heuristics(MazeSection start, MazeSection end) {

		    if (start != null && end != null) {
			    return (Int32)Math.Abs((start.ID.X - end.ID.X) + (start.ID.Y - end.ID.Y));
		    }
		    else {
			    return 99999;
		    }
	    }

        private class PathFindingNode
        {
            public MazeSection location;
            public PathFindingNode parent;
            public Int32 h;
            public Int32 g;
        }

        private MazeSection scatterNextDestination;
    }
}
