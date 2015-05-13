using Microsoft.Xna.Framework;
using mopacman.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Controllers
{
    class BasicIAController : GameComponent
    {
        public IControllable Player { get; private set; }

        public IControllable Target { get; private set; }

        public BasicIAController(MyGame g, IControllable player, IControllable target)
            : base(g)
        {
            this.Player = player;
            this.Target = target;
        }

        public override void Update(GameTime gameTime)
        {
            //Localizações dos elementos 
			var currLocation = this.Player.CurrentLocation;
			var currEndLocation = this.Target.CurrentLocation;
			var prevLocation = this.Player.PreviousLocation;

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
				List<MazeSection> opcoes = new List<MazeSection>{ currLocation.W, currLocation.N, currLocation.E, currLocation.S };
				
				foreach (var o in opcoes) {
										
					if (o != null) {

						//A seção precisa estar ativa e não ser a seção anteriormente visitada
						if (o.Allowed && (prevLocation == null || o.ID != prevLocation.ID))
						{
							//Seção não deve existir na lista de seções fechadas
							if(!closedSet.Any( x => x.location.ID == o.ID))
							{
								//...nem na lista de seções abertas
								if (!openSet.Any( x => x.location.ID == o.ID )) {
																		
									PathFindingNode node = new PathFindingNode();
									node.location = o;
									node.parent = currNode;
									node.h = this.Heuristics(node.location, currEndLocation); //heuristica manhattan
						
									/*
									//Se estiverem na mesma coluna, considerar linha
									if (currEndLocation->getID().x == o->getID().x) {
										node->g = (currEndLocation->getID() - o->getID()).y;
									}
									//Se estiverem na mesma linha, considerar coluna
									else if (currEndLocation->getID().y == o->getID().y) {
										node->g = (currEndLocation->getID() - o->getID()).x;
									}
									//Nos demais casos, usamos a heuristica de euclides
									else
									{
										node->g = std::sqrt(std::pow((o->getID().x - currEndLocation->getID().x), 2) + std::pow((o->getID().y - currEndLocation->getID().y), 2));
									}
									//*/

									//Ok, esse nó participará do teste 
									openSet.Add(node);
								}
							}
						}
					}
				}

				if (openSet.Count == 0) {
					break;
				}

				//Procuramos o melhor nó localizado (custo)
				var minNode = openSet[0];

				foreach(var m in openSet)
                {
					if (m.h < minNode.h) {
						minNode = m;
					}
				}
								
				currNode = minNode;
				currLocation = currNode.location;

				//Removendo da lista de nós abertos para pesquisa e incluíndo na lista de nós proibidos
				openSet.Remove(minNode);
				closedSet.Add(currNode);
				
			} while (currNode.location.ID != currEndLocation.ID);
									

			if (currNode.location.ID == currEndLocation.ID) {
				///*

				LinkedList<PathFindingNode> path    = new LinkedList<PathFindingNode>();
				PathFindingNode	lastPathNode		= currNode;
				MazeSection		nextSectionToMove	= null;
				MazeSection		mySection			= this.Player.CurrentLocation;
				
				while (lastPathNode != null) {
					path.AddFirst(lastPathNode);
					lastPathNode = lastPathNode.parent;
				}
								
				nextSectionToMove = path.First.Next.Value.location;
				
				if (nextSectionToMove.ID == mySection.N.ID) {
					this.Player.GoTo(EnumDirections.North);
				}
				else if (nextSectionToMove == mySection.S) {
					this.Player.GoTo(EnumDirections.South);
				}
				else if (nextSectionToMove == mySection.E) {
					this.Player.GoTo(EnumDirections.East);
				}
				else if (nextSectionToMove == mySection.W) {
					this.Player.GoTo(EnumDirections.West);
				}
				//*/
			}			
        }

        public Int32 Heuristics(MazeSection start, MazeSection end) {

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
    }
}
