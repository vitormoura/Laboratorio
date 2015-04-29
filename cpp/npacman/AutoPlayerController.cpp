#include "AutoPlayerController.h"
#include <vector>
#include <cmath>
#include "PathFindingNode.h"

namespace my {

	int heuristics(MazeSectionPtr start, MazeSectionPtr end) {
		if (start != nullptr && end != nullptr) {
			return std::abs((start->getID().x - end->getID().x) + (start->getID().y - end->getID().y));
		}
		else {
			return 99999;
		}
	}
		
	AutoPlayerController::AutoPlayerController(ControllablePtr self, ControllablePtr target) :
		m_self(self), m_target(target)
	{
	}

	AutoPlayerController::~AutoPlayerController()
	{
	}

	void AutoPlayerController::update(sf::Time t) {
		
		if (m_wait > 1)
		{

			//Localizações dos elementos 
			auto currLocation = m_self->getLocation();
			auto currEndLocation = m_target->getLocation();
			auto prevLocation = m_self->getPreviousLocation();

			//Essa é a seção inicial, 
			PathFindingNode* currNode = new PathFindingNode();
			currNode->h = 0;
			currNode->g = 0;
			currNode->parent = nullptr;
			currNode->location = currLocation;

			std::vector<PathFindingNode*> closedSet;
			std::vector<PathFindingNode*> openSet;
			std::vector<PathFindingNode*> toDelete;

			closedSet.push_back(currNode);
			toDelete.push_back(currNode);

			do
			{
				//Percorrendo as opções de caminho que temos
				std::vector<MazeSectionPtr> opcoes{ currLocation->W, currLocation->N, currLocation->E, currLocation->S };
				
				for (auto o : opcoes) {
										
					if (o != nullptr) {

						//A seção precisa estar ativa e não ser a seção anteriormente visitada
						if (o->allowed && o != prevLocation)
						{
							//Seção não deve existir na lista de seções fechadas
							if (std::find_if(closedSet.begin(), closedSet.end(), [&o](const PathFindingNode* n) { return n->location == o; }) == closedSet.end())
							{
								//...nem na lista de seções abertas
								if (std::find_if(openSet.begin(), openSet.end(), [&o](const PathFindingNode* n) { return n->location == o; }) == openSet.end()) {
																		
									PathFindingNode* node = new PathFindingNode();
									node->location = o;
									node->parent = currNode;
									node->h = heuristics(node->location, currEndLocation); //heuristica manhattan
																		
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
									openSet.push_back(node);
									toDelete.push_back(node);
								}
							}
						}
					}
				}

				if (openSet.size() == 0) {
					break;
				}

				//Procuramos o melhor nó localizado (custo)
				auto minNode = openSet.begin();

				for (auto m = openSet.begin(); m != openSet.end(); m++) {
					if ((*m)->h < (*minNode)->h) {
						minNode = m;
					}
				}
								
				currNode = *minNode;
				currLocation = currNode->location;

				//Removendo da lista de nós abertos para pesquisa e incluíndo na lista de nós proibidos
				openSet.erase(minNode);
				closedSet.push_back(currNode);
				
			} while (currNode->location != currEndLocation);
									

			if (currNode->location == currEndLocation) {
				///*

				std::list<PathFindingNode*> path;
				PathFindingNode*	lastPathNode		= currNode;
				MazeSectionPtr		nextSectionToMove	= nullptr;
				MazeSectionPtr		mySection			= m_self->getLocation();
				
				while (lastPathNode != nullptr) {
					path.push_front(lastPathNode);
					lastPathNode = lastPathNode->parent;
				}
								
				nextSectionToMove = (*(++path.begin()))->location;
				
				if (nextSectionToMove == mySection->N) {
					m_self->go(Directions::N);
				}
				else if (nextSectionToMove == mySection->S) {
					m_self->go(Directions::S);
				}
				else if (nextSectionToMove == mySection->E) {
					m_self->go(Directions::E);
				}
				else if (nextSectionToMove == mySection->W) {
					m_self->go(Directions::W);
				}
				//*/
			}

			//Limpando objetos criados durante o processamento
			///*
			for (auto& p : toDelete) {
				delete p;
			}
			//*/
						
			m_wait = 0;
		}
		else
		{
			m_wait +=  t.asSeconds();
		}
	}

	
}
