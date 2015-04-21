#include "GameScene.h"
#include "Game.h"

namespace my {

	GameScene::GameScene() : m_children(), m_parent(nullptr) {
		
	}

	GameScene::~GameScene() {

	}

	void GameScene::init() {
	}

	void GameScene::update(sf::Time t) {
		for (auto s : m_children) {
			s->update(t);
		}
	}

	void GameScene::render(sf::RenderTarget* t) {
		for (auto s : m_children) {
			s->render(t);
		}
	}
}